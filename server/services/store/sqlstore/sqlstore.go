package sqlstore

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"runtime"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/services/store"
	"github.com/mattermost/mattermost-plugin-api/cluster"

	"context"

	mmModel "github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

// SQLStore is a SQL database.
type SQLStore struct {
	db               *sql.DB
	dbType           string
	tablePrefix      string
	connectionString string
	isPlugin         bool
	isSingleUser     bool
	logger           mlog.LoggerIFace
	NewMutexFn       MutexFactory
	servicesAPI      servicesAPI
	isBinaryParam    bool
}

// MutexFactory is used by the store in plugin mode to generate
// a cluster mutex.
type MutexFactory func(name string) (*cluster.Mutex, error)

// New creates a new SQL implementation of the store.
func New(params Params) (*SQLStore, error) {
	if err := params.CheckValid(); err != nil {
		return nil, err
	}

	params.Logger.Info("connectDatabase", mlog.String("dbType", params.DBType))
	store := &SQLStore{
		// TODO: add replica DB support too.
		db:               params.DB,
		dbType:           params.DBType,
		tablePrefix:      params.TablePrefix,
		connectionString: params.ConnectionString,
		logger:           params.Logger,
		isPlugin:         params.IsPlugin,
		isSingleUser:     params.IsSingleUser,
		NewMutexFn:       params.NewMutexFn,
		servicesAPI:      params.ServicesAPI,
	}

	var err error
	store.isBinaryParam, err = store.computeBinaryParam()
	if err != nil {
		params.Logger.Error(`Cannot compute binary parameter`, mlog.Err(err))
		return nil, err
	}

	err = store.Migrate()
	if err != nil {
		params.Logger.Error(`Table creation / migration failed`, mlog.Err(err))

		return nil, err
	}
	return store, nil
}

// computeBinaryParam returns whether the data source uses binary_parameters
// when using Postgres.
func (s *SQLStore) computeBinaryParam() (bool, error) {
	if s.dbType != model.PostgresDBType {
		return false, nil
	}

	url, err := url.Parse(s.connectionString)
	if err != nil {
		return false, err
	}
	return url.Query().Get("binary_parameters") == "yes", nil
}

// Shutdown close the connection with the store.
func (s *SQLStore) Shutdown() error {
	return s.db.Close()
}

// DBHandle returns the raw sql.DB handle.
// It is used by the mattermostauthlayer to run their own
// raw SQL queries.
func (s *SQLStore) DBHandle() *sql.DB {
	return s.db
}

var ConflictError = errors.New("conflict")

func SimpleOptimisticRetryableImmTx(s *SQLStore, ctx context.Context, f func(tx *sql.Tx) error) error {
	_, _, err := OptimisticRetryableImmTx(s, ctx, func(tx *sql.Tx) (*struct{}, error) {
		return nil, f(tx)
	})
	return err
}

func OptimisticRetryableImmTx[R any](s *SQLStore, ctx context.Context, f func(tx *sql.Tx) (*R, error)) (*R, string, error) {
	for {
		r, version, err := OptimisticImmTx(s, ctx, f)
		if err == nil {
			return r, version, nil
		}
		if err != ConflictError {
			return nil, "", err
		}

		s.logger.Info("retryTx", mlog.Err(err))

		durMs, err := rand.Int(rand.Reader, big.NewInt(500))
		if err != nil {
			return nil, "", err
		}
		timer := time.NewTimer(time.Duration(durMs.Int64()+100) * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, "", ctx.Err()
		case <-timer.C:
		}
	}
}

func OptimisticImmTx[R any](s *SQLStore, ctx context.Context, f func(tx *sql.Tx) (*R, error)) (*R, string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	_, err := s.db.Exec("select mv_commitgroup_begin(); select mv_commitgroup_lock_disable()")
	if err != nil {
		return nil, "", err
	}
	rollback := true
	defer func() {
		if rollback {
			_, err := s.db.Exec("select mv_commitgroup_rollback()")
			if err != nil {
				panic(err)
			}
		}
	}()

	txn, err := s.BeginImmTx(ctx, nil)
	if err != nil {
		return nil, "", err
	}

	r, err := f(txn)
	if err != nil {
		if err := txn.Rollback(); err != nil {
			s.logger.Error("failed to rollback transaction", mlog.Err(err))
		}
		return nil, "", err
	}

	if err := txn.Commit(); err != nil {
		return nil, "", err
	}

	rollback = false
	row := s.db.QueryRow("select mv_commitgroup_commit()")
	var maybeVersion sql.NullString
	err = row.Scan(&maybeVersion)
	if err != nil {
		panic(err) // We can't recover from this.
	}

	if !maybeVersion.Valid {
		// nothing changed
		return r, "", nil
	}

	if maybeVersion.String == "conflict" {
		return nil, "", ConflictError
	}
	return r, maybeVersion.String, nil
}

func (s *SQLStore) BeginImmTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := s.db.BeginTx(ctx, opts)
	if err == nil {
		_, err = tx.Exec("ROLLBACK; BEGIN IMMEDIATE")
	}
	return tx, err
}

// DBType returns the DB driver used for the store.
func (s *SQLStore) DBType() string {
	return s.dbType
}

func (s *SQLStore) getQueryBuilder(db sq.BaseRunner) sq.StatementBuilderType {
	builder := sq.StatementBuilder
	if s.dbType == model.PostgresDBType || s.dbType == model.SqliteDBType {
		builder = builder.PlaceholderFormat(sq.Dollar)
	}

	return builder.RunWith(db)
}

func (s *SQLStore) escapeField(fieldName string) string {
	if s.dbType == model.MysqlDBType {
		return "`" + fieldName + "`"
	}
	if s.dbType == model.PostgresDBType || s.dbType == model.SqliteDBType {
		return "\"" + fieldName + "\""
	}
	return fieldName
}

func (s *SQLStore) concatenationSelector(field string, delimiter string) string {
	if s.dbType == model.SqliteDBType {
		return fmt.Sprintf("group_concat(%s)", field)
	}
	if s.dbType == model.PostgresDBType {
		return fmt.Sprintf("string_agg(%s, '%s')", field, delimiter)
	}
	if s.dbType == model.MysqlDBType {
		return fmt.Sprintf("GROUP_CONCAT(%s SEPARATOR '%s')", field, delimiter)
	}
	return ""
}

func (s *SQLStore) elementInColumn(column string) string {
	if s.dbType == model.SqliteDBType || s.dbType == model.MysqlDBType {
		return fmt.Sprintf("instr(%s, ?) > 0", column)
	}
	if s.dbType == model.PostgresDBType {
		return fmt.Sprintf("position(? in %s) > 0", column)
	}
	return ""
}

func (s *SQLStore) getLicense(db sq.BaseRunner) *mmModel.License {
	return nil
}

func (s *SQLStore) getCloudLimits(db sq.BaseRunner) (*mmModel.ProductLimits, error) {
	return nil, nil
}

func (s *SQLStore) searchUserChannels(db sq.BaseRunner, teamID, userID, query string) ([]*mmModel.Channel, error) {
	return nil, store.NewNotSupportedError("search user channels not supported on standalone mode")
}

func (s *SQLStore) getChannel(db sq.BaseRunner, teamID, channel string) (*mmModel.Channel, error) {
	return nil, store.NewNotSupportedError("get channel not supported on standalone mode")
}
