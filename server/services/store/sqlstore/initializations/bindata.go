// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/templates.json
package initializations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\xcd\x8e\x5c\xc7\x75\xff\xab\x5c\xdc\xcd\x7f\x33\x07\xff\xfa\x38\x55\x75\x6a\x76\x92\x68\x29\x04\x14\x8b\x96\xe4\x7c\xc0\x10\x8c\x53\x55\xa7\xc8\xb6\x9a\xdd\x83\xee\x1e\xd2\x03\x82\xab\x6c\xed\xc8\x31\x04\x18\x89\x8c\xd0\x88\x91\xc0\x4b\x07\xc9\x22\x08\x90\x8d\x1f\x45\x2f\x10\x3e\x42\x50\x3d\x1f\x1c\xce\xdc\x1e\x16\x47\x24\x67\xec\x04\xe0\xa2\x79\xef\xed\x73\xab\xcf\xe7\xef\xfc\xea\xd4\x3c\x19\x1f\xc9\x6a\x3d\x5b\x2e\xc6\x7d\xbd\x37\x16\xde\xc8\xb8\xaf\xbd\x0a\x41\x3b\x6f\x5c\x08\xb4\x37\xa6\xf9\x32\x7f\xb9\x1e\xf7\x7f\xf4\x64\x9c\x95\x71\x7f\x54\x2a\x64\xce\xde\x03\x72\x72\x80\xc9\x29\x20\x97\x0c\x18\x15\x4b\x14\xe7\x38\xab\x3c\xee\x8d\x07\xbc\x92\xc5\xe6\x6e\xfb\x86\x49\x29\x94\x6c\x0b\x64\xeb\x19\x50\x94\x02\x12\x55\x81\x7d\xb1\x9a\xb3\x72\x36\x87\x71\x6f\x5c\x2d\x97\xaf\xf5\xfc\x3a\x3f\x90\x87\xbc\x5d\xf8\xe6\xe8\x40\xc6\xfd\x31\xf3\xaa\x8c\x7b\xe3\x66\xb6\x99\xb7\xff\x7e\x26\x9b\xcd\x6c\x71\x7f\x3d\xfc\xf0\xaf\xc6\xbd\xb1\xce\x64\x5e\xd6\xe3\xfe\x93\x71\x96\xdb\xef\x1d\x9f\x3f\xfb\xdb\x6f\xfe\xfb\x3f\xbe\x6a\x6b\x5d\x2d\x0f\x64\xb5\x99\xc9\xf6\x36\xc7\x60\x4a\x0e\x0c\xae\x62\x06\x74\xc5\x00\x29\x44\xa0\x6c\x28\xfb\x18\x74\xa8\x7a\xdc\x1f\x4b\x09\xc9\x72\x88\x20\xc9\x14\xc0\x68\x1d\x50\x74\x11\x4c\x54\x2e\x72\x89\x2e\x58\x1e\xf7\xc6\x62\x8b\x27\x93\x2a\x88\x0a\x08\x18\x4b\x04\x2a\xd5\x41\xb0\x46\x45\xa3\xb3\x29\xc6\x8e\xfb\x23\x85\xea\x62\x20\x84\x44\x2e\x02\x66\x83\x40\x92\x04\x74\xc8\xc1\x7b\x51\xa4\x4b\x19\x9f\x3e\xdd\x1b\xf3\x4a\x78\x23\xef\x6d\x8e\xed\xe4\x8d\x8e\xd1\x69\x34\x7b\xe3\xe1\x41\x79\xe9\x86\x51\xe8\xa2\x8a\x7b\x63\x91\xb9\x1c\xdf\x50\x4f\xf7\x4e\x8d\x18\x12\x3b\xb6\x0a\xac\x44\x0f\x28\x52\x80\x0a\x57\x20\x87\x39\x29\xe6\x60\x55\x78\xd9\x88\xd1\x62\x8d\x26\x05\x88\xb5\x68\xc0\x14\x2d\x50\xcd\x1a\x14\x22\xa7\xc2\x92\x73\x8e\xe7\x8d\xd8\xf9\xfc\x65\x23\x3e\x9a\xc9\xe3\x73\x46\x7c\x7f\xc9\xab\x32\xfc\xc5\xf1\xc5\x17\x36\x6c\xa6\xfe\x64\x55\x64\x35\xee\xff\xe8\x8b\xbd\x31\x2f\xe7\x87\x0f\x17\x7f\x39\x2b\x9b\x07\xed\xf6\xd3\xf6\xe8\x7c\xd3\xee\x3e\x39\xf9\xb4\x3e\x7e\xb0\x19\x9a\x37\x5b\x87\x1f\x79\x51\xc6\xa7\x7b\xe3\x83\x59\x29\xb2\xf8\xe4\xa0\x5d\xbd\x5b\x4e\x9e\x5b\x2f\x57\x9b\xe3\x4b\x27\x17\xda\xb2\x3e\x3f\x5e\x61\x5a\x1e\xfb\xd9\xa3\xd9\x7a\x96\xe6\x72\xe1\x9b\x27\x57\xef\x1d\xfb\xd4\xd1\xc9\xf5\x09\xcb\xa9\x68\x9d\xd3\xe1\xb2\xe5\xce\x6e\x4c\x5a\x4e\x24\xba\x6a\x23\x64\x31\x19\xb0\x78\x02\xd2\xa4\x41\xe7\xa4\x8a\x91\xca\x5e\xd7\x1b\xb2\xdc\x85\xf0\xfb\x88\x57\x45\x16\xb3\xc5\xfd\xe9\xe0\xfb\xd9\xbf\x5d\x8a\xbc\x12\x42\x48\x6c\x13\x50\x30\x04\xa8\x8a\x06\x0a\xec\xc1\x45\x54\x3e\xa5\x94\x6a\x52\xe3\xfe\x68\x51\x52\xcc\xc6\x41\x71\xa9\x9e\x84\x94\x8b\x02\x35\x60\x6d\xd9\xc2\x2a\x09\xd3\xb1\x62\x51\x29\xe3\x2f\x6b\x5c\x5b\x47\x21\xaa\x69\x8d\xeb\x6c\x9d\xaf\x4d\xbe\x21\x06\x4c\x5a\x80\x7c\x8b\xe0\x8a\x12\x0b\x72\xd0\x16\x6f\x47\xc2\xfb\x9c\xd7\x5f\x4e\x29\x7b\xdc\x1b\x67\xeb\xcf\xe5\xe1\xc1\x7c\x9b\xe0\x37\xab\x43\xb9\x66\xd2\x43\x0c\x92\x2b\x6a\x28\xd5\x15\x40\x23\x1a\x28\x72\x82\xa0\x51\x7c\x70\xc2\x3e\xe8\x77\x90\xf4\x0c\x2a\x67\x70\xc2\x90\x06\xb5\x56\xc6\xee\x32\x64\x41\x8b\x49\x41\x70\xec\x01\x83\x2d\x40\xc6\x0a\x44\x47\x9c\x23\x71\xf5\xbe\xdc\x90\x21\x2f\x26\xbd\xa3\xe1\xde\x6a\xb6\x5c\xcd\x36\x47\x6f\x2b\xeb\xdd\x5f\x2d\x0f\x0f\xde\x3f\xda\x2e\xbb\xcb\x5c\x53\x89\x72\xf4\x4e\x5c\x8e\x6c\xc1\x26\x66\xc0\xaa\x03\x90\xf6\x19\x0c\x27\x93\x38\xd8\x22\x25\x8e\x6f\x3e\x9d\xf6\xad\x78\x2a\xe9\xea\xe0\xb5\xb3\x76\xc2\x73\x8c\xb3\xde\xb8\x69\xcf\x31\x4a\xab\x84\x48\x90\x4d\x34\xcd\xed\x13\x10\xa7\x00\xb9\x78\xe7\x75\x4e\x59\xa5\x0b\x29\xc0\x27\xb1\x31\x67\x0d\xd5\xb8\x0c\x18\x52\x05\x2a\x18\xa1\x7a\x41\x26\x95\x2b\x05\x73\xde\x73\x3a\x9f\x7f\x35\xe6\xd9\xf0\x6a\x33\xf0\x50\x78\x36\x3f\x1a\x7e\xb2\x3c\x5c\x2d\x78\x3e\x95\x13\xbe\xfd\xf5\xcf\x27\xc1\x4f\xf5\x35\xa7\x44\x90\xd9\x79\xc0\x14\x2c\x90\x95\x04\x36\xa0\x0d\x29\xb2\xf7\xa1\x2d\x37\x55\x67\x52\x15\x0f\x9c\x5b\xba\x88\x48\x40\x46\xd7\xf6\x2c\xb3\xce\x21\x29\xdc\x82\x1f\x9f\xb4\xc1\x98\x20\x65\x4d\x80\xae\x66\x20\x8a\x02\x49\x32\x52\xcd\x42\x4a\x6a\xab\x69\x4c\x26\x86\x50\xa1\x09\x05\x74\x21\x01\xa1\x4e\x20\x26\x85\x1a\x7c\x4d\xce\xb8\x26\x2d\x06\xe3\x8a\x46\x28\x8e\x09\x90\xc4\x01\x15\xd1\xe0\x2b\x65\xa5\x90\xa3\xa7\x6d\x79\x30\xe8\xd8\x9a\x02\xd5\x13\x01\xfa\x6d\x3d\xa9\x2d\xb5\x04\x8d\x26\xeb\x64\xe3\x54\x79\x08\xda\x39\x17\x10\xf5\x45\xdf\x68\x37\x48\x7b\x4d\x3b\x7c\xc3\xe7\xcc\xc4\x19\x72\x42\x02\x64\x1b\x80\xc8\x21\x10\x26\x96\x98\x74\x24\xa4\x1b\xf2\x8d\xcb\x59\x65\xbd\xe1\xcd\xe1\xfa\x5d\xe4\x94\x2e\x47\x7a\x37\xe0\xab\xcf\x71\x3a\xbd\xf5\x72\x52\x09\xda\xb5\xe4\x11\x2e\x61\xf0\xe6\x38\x9a\x8c\xdf\x81\xe4\x3a\x8b\xc5\x39\xc7\x79\x23\xa5\xe6\x54\x81\xa7\x5e\x71\x6f\xb5\xfc\x89\xe4\xcd\xd0\xc0\xc3\x65\xcf\xb8\x77\x2e\x3b\x9c\x76\x80\x5d\x58\x61\x6f\x5c\xf0\xc3\x93\x8c\x74\xec\x72\xcb\x33\x8b\x3e\x69\x5e\xb6\x5c\x8d\xfb\xdb\xe4\xf3\x41\xfb\x7c\x47\x2a\x1f\xce\x37\x0d\xa9\x94\x7e\x9c\xf1\x88\xe7\x87\xed\x1d\xdf\x97\x9f\x6e\x86\x1f\x1e\x8c\x4d\xb7\x97\x44\xff\xb5\xcc\xe7\xcb\xc7\xa7\x92\x3b\xdb\xb6\x53\xc9\x77\x17\xc3\xbd\xd5\xf2\xfe\x4a\xd6\xeb\x69\xe9\x1f\xad\x44\x16\xe7\x84\xb7\xc2\x57\x20\xe6\xe0\x01\x8b\x44\x20\x55\x09\x9c\x0e\xa2\x7d\x2d\x9a\xdc\xf9\x65\x7f\xb0\x7c\x78\xd0\xfc\xa2\xfc\xe1\x37\x7f\xf8\xcd\xf3\x67\x7f\xff\xb3\xe9\x57\xbc\xbf\x5a\x3e\x3e\x7b\x45\x57\xc5\x7d\xf1\x8a\xf7\x56\xf9\xc1\xec\x91\x94\xf1\xe9\x17\x67\x1e\xb0\x96\xb9\xe4\xcd\x78\xe6\x88\x9d\xe5\xff\xc4\x9e\xe7\x80\xc9\x95\x16\xfd\x54\xca\x99\x5a\x6c\xaa\xc9\xa9\x0a\xd5\xf9\x08\x98\xb8\x00\xb1\x65\x28\x45\xbb\x6c\xab\x57\x5a\x9d\x57\xcb\x9f\xcd\xee\x3f\x18\x9e\x3f\xfb\xfa\x9f\x7b\xec\xd9\x85\x21\x5f\xc8\xfe\x73\x29\xb3\xc3\x87\xbb\x4c\xc9\x47\xa7\x62\x23\xb1\x0b\xde\x04\x50\xb5\x59\x32\xe8\x02\xe4\x54\x81\xa8\xab\x95\x12\x6b\x31\xda\x9e\x13\xfb\xf1\xf2\xf1\x55\x1a\x36\xec\x0a\x5b\xa3\x20\xd8\x96\xc1\x55\x6b\xa9\x02\x85\xa6\x0f\xd1\x2e\x97\x2a\x5b\x65\x9d\x68\xf8\x0e\x6f\x64\xf8\x60\x9b\x63\xca\x4b\x5a\x7e\xf1\x82\xe3\x0c\x54\x3e\x9f\x3d\x94\xed\x7b\x8b\xac\xf3\x6a\x76\x70\x92\x8e\xdb\x8f\x38\xe3\x38\x7e\x39\x81\xfe\x77\x02\xa3\x29\x1e\xc1\x90\xc1\xe8\xa7\x73\x98\xa5\x8a\x49\x54\x80\xaa\x2b\x37\x1c\x90\x81\xb2\x76\xc0\x51\x05\xa5\x7c\x8d\xec\xdc\xed\x00\x46\x1f\x0b\xaf\x16\xc3\x66\x39\x1c\xf0\x6c\xb1\xd9\xc1\x07\xfd\xee\xba\x78\x28\x84\xec\x6c\xe4\x0a\x56\xc5\xd6\x95\x27\x0d\x64\x4d\x84\x62\x94\xd4\x20\x51\x38\x97\x7e\x3c\x14\x59\x45\x25\xd6\x42\xd2\x5a\x01\x9a\xd6\xe3\x47\x15\x41\x5b\x93\x9d\x52\x26\x47\x25\xef\x02\x0f\xb5\xf8\x77\x97\xb0\x72\xd0\x0e\xa3\xb6\x84\xd3\x2e\x81\x31\x56\x47\x92\x21\xa1\x36\x80\x99\x5a\x61\xa8\x08\x31\xb0\xf1\x9e\x62\x20\xbe\xd0\x65\x95\x50\xc9\x06\x2a\x40\x1a\x5d\xfb\x86\x01\xc2\xe2\x81\x6b\x4d\x92\x42\x2d\xa1\xf0\x79\x97\xe8\x7c\xfe\xd5\x78\xe8\xf0\xfe\x7a\x78\xfe\xec\x17\xff\xf8\x1a\x78\x68\xfc\xf1\x8f\x4f\xbe\x6e\x48\xed\x00\x47\x2d\xb5\x2c\xca\xec\x24\x18\x67\x8b\x3c\x3f\x2c\xb2\x7e\xe1\x5a\xc7\x00\xc9\xa8\xa0\x03\x17\x0b\x2e\xa0\x06\x44\xed\x81\x6c\x6d\xd6\xd3\xd6\xd6\x5a\x6d\xd1\x67\x49\x71\x8b\x62\x74\x2d\x29\x39\xed\x40\x4a\x31\x80\x5c\x5d\x2b\x2e\x19\xd0\x06\x29\xc6\x68\xcd\x18\xc7\x2f\x9e\x5e\x9b\xe4\x7a\xf2\xf2\xfa\x6a\xa8\x96\x0b\x1a\x48\x56\x33\x20\x67\x03\xa4\xab\x02\x43\xd2\xe2\xce\x25\x8b\xcd\x09\x57\xf2\x48\x56\x6b\x29\xe3\x7e\xe5\xf9\x5a\x9e\xbe\x8c\xd6\x36\x9c\xe6\xf2\x9a\x68\xcd\x29\xad\x43\x71\x06\x52\xce\x01\x30\x38\x05\x64\x98\x81\xac\x66\xeb\x74\x46\xe4\x86\xd6\x3a\xf5\xe7\x55\x24\x57\xd1\x83\x95\xf6\x18\x79\x69\x9d\x8a\x05\x93\x22\x05\x44\x25\xac\x73\xb3\x7f\xcf\xaf\x9d\x6a\x28\x11\x8d\x33\x34\xc1\xe2\x21\x06\xe3\x77\xb1\x78\x4e\x69\x9b\x89\x5b\xab\x20\xad\xe4\xaa\x0a\x14\xad\x01\xad\x23\x46\xc5\x5a\xab\x72\x4b\x1a\xca\x4f\x0e\x64\x31\xac\x64\x33\x5b\xc9\x43\x59\x6c\x06\xce\x79\x79\xb8\x2b\x81\x7e\xf5\x2f\x7f\x8a\x0d\x25\xa9\x24\x5b\x8c\x55\x1b\xe8\xc5\x68\x08\x08\x29\x02\x25\x65\x34\xa9\x8a\xd1\xe1\x8e\x04\x4a\xde\x39\x37\xd5\x17\x28\xa5\x70\x47\x4d\x75\xc6\xa2\x36\x59\x83\x75\x96\x01\x4d\x4d\x40\x9a\x2a\x70\x49\x3e\x19\x71\x29\x87\x72\x3b\x18\xde\xef\xcb\xe3\xe1\x3b\x70\x8e\x6f\x9d\xee\x45\xf4\xda\x5e\xea\xe7\xdb\x0d\xa7\x95\xdf\x41\xf7\x3a\x22\xa7\x89\x04\x24\x04\x03\xe8\x5b\x68\x06\x36\x40\xde\x2a\xca\x35\x63\xc8\xea\x76\x84\xe6\xa7\x87\x8b\xc1\x0e\x9b\xd9\x43\x59\x0f\x3c\x3c\x16\x99\x34\xc4\xf3\x67\x5f\xfd\xcd\xed\x88\x4a\x2f\x41\xdb\x28\x08\xce\xba\x06\x44\x52\x02\xca\x4a\x43\x48\xca\x44\x76\x21\x91\xe2\xfe\xa8\xac\x35\xd9\x1a\x9d\x86\x84\xa1\x02\xea\xe6\x47\x55\x17\x30\x96\xbc\xf7\xc1\x90\x52\x34\x1d\x95\x68\x8d\xbe\x4c\x1e\x6f\xa3\xd2\x92\xde\x01\x6b\x9c\xf8\xaa\x5d\xd2\x60\x4c\x22\x40\x69\xd5\xd7\x18\x06\x2e\xc6\xd8\xd2\xba\x24\xa5\x6e\x08\xd6\x5c\x74\x0b\x69\x05\x78\x78\xef\xde\xdd\xa1\xc8\x7a\x76\x7f\x31\xed\x15\xdf\xfc\x76\x8a\xff\xeb\xaa\xaa\xfb\x23\x1a\x14\x76\x62\x21\x32\x6b\x40\x15\x1c\x50\x76\x19\x94\x4e\x88\xc9\x7b\xf1\xdb\x3d\x91\xae\x52\xbe\x3f\x52\x76\x2e\x54\x1f\x21\x29\x29\x80\x5e\x32\x90\x65\x0b\x24\x95\x53\x2c\x18\xa4\xba\xde\x52\xbe\x3f\x6a\xa4\x68\x2c\x29\xd0\x6c\x0d\xa0\x29\x05\x48\x59\x04\x56\x59\xd8\x9a\x94\x83\xf8\xde\x8a\xbf\x3f\x8a\xe7\x50\x4d\x0c\x80\x88\x0a\x30\x90\x05\xe2\x64\xc1\x72\x75\x89\xbd\xc9\xac\xf5\xf4\x1e\x85\x8b\x51\x11\x5d\xce\x3e\x56\x2b\x72\x66\x87\x9b\x75\xe6\x86\x1d\xa4\xd0\xb5\x33\xcb\x25\x52\x48\x56\xeb\xe5\x82\xe7\xc3\x47\x4b\x9e\x77\xb2\x42\x7d\x84\xdf\x6b\xb1\x42\xe7\x38\x84\xce\x54\x74\xda\x90\x7f\xbe\x1c\xee\x2c\x7b\xf8\x83\xce\xce\xed\x54\xee\x9d\xe5\x6c\x71\xbf\x83\x09\x8a\x94\x8a\xb0\x09\xa0\x72\x16\xc0\x5c\x04\xa8\x64\x0f\x46\x1b\x2e\xc5\xb6\x00\xe1\x97\xc4\x2e\x64\x38\x66\x80\xae\xe0\x69\xfa\x68\xcc\x13\x0d\x7f\xc0\x1b\xb9\xbf\x5c\xbd\x92\xa7\xb9\x77\xb8\x3a\xd8\x62\xf4\xe3\x86\xbe\xa7\x59\x3c\xc7\x7b\xcc\xaa\x0c\x9f\x7d\x39\x9b\xcf\x7b\xe8\xb1\x2e\x20\xf5\x42\xf8\x87\xb3\x05\x2f\xb2\x4c\x0b\xfe\x64\xc5\x8b\xfb\x67\xcb\xee\x2a\x06\xe7\x18\x26\xe1\xf9\xe6\xc1\x95\xba\xee\x29\x69\x2f\x18\x9b\x43\x19\xee\x34\x9c\xf3\x7a\x2c\x67\x67\x9f\x7f\xba\xea\x1f\xe8\x69\x55\x5c\x90\xda\x09\x7e\xcf\xa4\x9a\x2e\xa9\x9d\xc5\xfb\x4c\xaa\xed\x92\x5a\x9c\x0d\x3a\x7b\x0b\xde\xb7\xb5\x62\xd3\x40\xb0\x04\x19\x4b\xc6\x24\x8c\x84\xe7\x69\xb6\x1f\xe0\x94\xcd\xae\xe0\xbf\xbe\xfd\xe6\xf7\xc7\x55\xee\x95\x0c\x58\xd0\xce\x68\x32\x41\x4d\xe0\x02\xbf\xfd\xb7\x23\x61\x57\xeb\x8a\x4d\x16\x54\x2a\x01\x90\xd5\x56\x2f\x08\xe8\x48\x45\x93\xad\xaf\x64\x6f\xc7\x74\xc0\x1d\xde\x70\xe2\xb5\x0c\x9f\x1d\x3f\x39\x89\x0a\x7e\xf9\x5f\xd7\x9d\x87\x7a\xb3\xa3\x01\x5d\x94\xf1\x24\xe8\x27\x8c\xc1\x4e\xf4\xe3\x3a\x62\xf0\x7e\x47\xd9\x0d\x14\x13\x59\x4c\x80\xda\x38\x40\x97\xf5\xf1\x3c\x54\xb1\x3e\xf9\x98\xc5\x90\x71\xb7\x84\xb4\x3a\x1a\x3e\x3b\x58\x5d\xa4\x30\xdf\xd2\x26\x5e\x27\x79\x72\x43\xe4\xd2\x75\xb6\x02\x3b\x59\xa3\x6b\xd3\x41\x56\x29\xe7\xec\x24\xea\x43\xb2\x7e\xc7\x1e\x72\x28\x4a\xa3\xa7\x08\x52\x63\x05\xcc\x35\x00\x79\x2a\x90\xc4\x99\x20\xbe\x5a\x5f\xfc\xed\x68\xf9\x3f\x14\x29\xc3\x87\xf3\xc3\x5a\x8f\xa6\x13\xc8\x2f\xfe\xf5\xdd\x8f\x75\x19\x45\x1e\xa7\x42\xde\x44\x72\x6e\x47\x9f\x1f\xad\x51\xc5\xa3\x82\x5c\x5b\xe2\x76\xd5\xb6\x8c\xa6\xa1\xb0\x32\x4a\x07\xc9\xca\xa5\xdb\xd1\xd0\xdd\xcd\xcb\xc5\x7a\x28\xcb\xc5\xff\xdb\x0c\x65\xb6\x3e\x98\xf3\x0e\xdd\xff\xfc\x57\xd7\xed\xe7\xba\xb8\xe5\x1b\xea\xe7\xa4\x20\x3b\xdb\x9a\x2f\xaf\x9a\x34\x9d\x80\xc8\x12\x98\xec\xbc\x90\x88\x8b\x55\xfa\xfb\xb9\x9c\x48\x72\xe1\x0c\x96\x92\x00\x16\xeb\x81\x82\x36\x0d\x71\x38\xca\x99\x98\x24\x4e\xf7\x73\xd1\x92\x51\x13\x6c\x92\xd5\xca\x38\xb5\x03\x1e\x74\xc6\xdd\x8e\x7e\xee\xda\x51\xbb\xb3\x9f\xeb\xdf\xe5\xef\x8a\xda\xd7\xec\xe7\xee\xcd\x16\x5f\x9e\x75\x1a\x3d\xe1\x7e\x8d\x86\xae\xd8\xc0\x5e\x57\x84\x6a\x5b\xfb\x5f\x52\xeb\x2f\x4c\x01\x27\x8a\xa3\x66\x36\x4a\xae\xd3\xd0\x15\x4e\x5c\x4a\x4c\xc0\xa5\x6a\xc0\x12\x2b\x50\x50\x0d\x0a\x39\xce\x3e\x4b\xf5\xca\xf4\x36\x74\x57\x01\xd6\x5f\x7f\xdd\x09\x58\x4f\xe7\x84\x27\xc7\x59\xa3\x73\xbb\x66\xd9\x62\x40\xe7\x55\x8a\x90\xab\x6b\x90\x5b\xe5\xd6\xe1\x29\x30\x0d\xe3\xa6\xa0\xad\xd1\x72\x3b\x6a\xcd\x47\xcb\xa1\x2e\x57\x03\x0f\x8f\x79\xbe\x83\xda\xfc\xbb\xdf\xde\xc0\x10\xb1\xf6\x2e\x5c\x22\xf5\xdb\x0d\xeb\x8d\x33\xd3\x5a\xe7\xe4\xbc\x15\x4e\x90\x50\x05\x40\x8b\xdc\xd2\x0e\x81\xa6\x92\x89\x82\x28\x65\xea\xad\x99\x12\x2b\x87\x32\x94\xe3\x76\xf6\xed\xcf\x9e\xf6\x35\xda\xb7\x67\x4e\x6c\x72\x00\xcc\x19\x13\xa7\x27\x07\x2d\x59\xb7\x63\x13\x30\x11\x45\x22\xcf\x50\x2b\x45\x40\xb1\xa5\x15\x5d\x0f\xce\x44\xf1\x41\xb3\xad\x91\x6e\x09\x02\x79\x78\xb0\x5c\x6d\x86\xff\x3f\x7c\xef\xa7\xed\xc3\x74\x2c\xfe\xc3\x3f\x5d\x17\x7d\x78\x11\x8e\x3e\x47\xc0\xec\x35\x60\xf4\x04\xe4\x1c\x02\xa6\xe0\x6c\x10\xaa\x61\x3b\x42\xd9\x89\x3e\x24\xfb\x62\x52\x76\x50\xaa\x49\x80\x35\x04\x20\x0c\x11\xc4\xc5\x2c\x49\xd9\x18\xd1\xf7\xa3\x8f\xac\xb4\x0f\x3e\x33\x24\xde\xb2\x29\x16\x81\x92\x38\xc8\x59\x8c\xa1\x50\x98\x7d\xf7\xfe\xf1\x77\x61\x93\xdb\x4f\xb0\x71\x12\x7d\xa0\x0b\x3b\x30\x6e\x8a\x45\xd7\xac\x05\x48\x69\x0d\xd8\x9c\x8b\x74\x4e\xc0\x5e\x25\xa5\x6d\xcc\x89\x6e\xc9\xc4\xfb\x7b\xf3\x5d\xd8\xe4\x8a\x59\x8c\xae\xb1\xaa\x7d\x1d\xe2\xde\x4b\x63\x1b\x7b\x9d\x94\x86\x36\xa6\x97\xaf\xd0\x7a\xd7\x30\xc8\x1b\x3a\x73\x74\x9d\x41\x8a\xce\x39\xcd\xce\xf1\xbf\x2e\x65\x4f\xb5\xc4\xc6\x78\x8c\x7e\x6a\xe4\xde\x7a\x6d\x77\x15\xcc\x94\x72\xf1\x2e\x30\xf0\xd6\x75\x7d\x09\x0d\x6f\x39\xc8\x1c\x91\x30\x98\xaa\x24\xde\x0e\x46\x66\xeb\xb6\xc3\xf3\x67\x5f\xff\xee\x8f\x63\x8e\xa8\x2b\xdb\xfe\xdf\x1c\xd1\x3b\x9f\x23\xa2\x88\x76\x02\x56\xa2\x53\x14\xb4\xde\x11\x25\xd9\x27\xc6\x6a\x01\x5d\xd8\x12\x47\x05\x88\x63\x84\x5a\x7d\x52\xc1\x15\x56\x7c\x53\x24\xc6\x85\x28\xf9\xde\xc1\x2c\xaf\x87\x6f\xbf\xf9\xfd\x1f\x47\x90\x74\x6d\x70\xbf\xb9\x20\xe9\xf4\xae\xff\xed\x41\x62\xb5\xb7\x01\xa7\x82\x04\x03\x45\xbb\x83\x5d\xcd\x2a\x17\x69\xdd\x85\xa1\x68\x01\xa3\x44\x20\x94\x0c\xce\xb8\x92\x14\x3b\x2c\x36\xdd\x0e\x14\xf4\xfe\xd1\xf0\xd9\xbb\x3b\xa1\xd3\x07\x0f\x6e\xea\x78\xb4\x69\x46\x8d\x38\x79\x52\x8f\x02\xee\xe8\xa9\x32\x55\x74\xa4\x0a\x58\x97\x14\xa0\xb5\xdb\xee\x3f\x43\x26\xed\x29\x3b\x26\x7f\x91\x6f\xbb\xc9\x8d\x9c\xab\x6d\xdd\x4b\x51\x77\xf6\x91\x5d\x13\x4c\x6f\xd0\xbf\x3a\x53\xcb\x9f\xde\xe6\xd1\xdb\xce\x82\xc6\x19\x37\x39\xd7\x68\xb5\xc1\xa0\x76\xf4\x82\x9d\x7e\xbb\x83\x89\xbe\xb6\xd7\x5f\x64\xa2\x3f\x5d\x72\x79\xc8\x07\x5d\x14\x74\xa7\x07\x7d\x97\x83\x66\x9d\x5b\x15\x67\x07\xcd\x96\x9b\x61\x7b\xc4\x56\x4a\x0f\x17\xdd\x49\x45\x5c\xef\xb0\x19\x61\x0c\xde\x27\x06\xe7\xd9\x01\x52\xd1\x40\xe4\x2b\x18\xed\x83\xa9\xa8\x9d\x8d\x6e\xe2\xb0\xd9\x2b\xc7\x8c\x3a\xbd\xfc\x44\xeb\xdb\x70\x7a\x85\xce\x5f\x56\x4a\xe7\xec\xe0\xe9\xba\x1b\x70\xdc\xe2\xc6\x57\x6b\xa4\x93\x47\x3a\xdb\x52\xe0\xf5\x97\xc7\x7d\xdb\xa4\xe8\x73\xc3\x67\x9d\xdb\x63\xa7\x82\xdf\x3f\xbc\x7f\x7c\xae\xe4\x0a\x2d\x77\x26\x89\x53\xdf\x3e\xdd\xf2\x7f\x1d\xdf\xee\xa4\xae\x4e\x57\x7d\xfc\x8e\xa1\x6f\xc8\xa8\x73\x53\xee\x82\xec\xbe\x51\xa3\xce\x01\xce\x0b\xb2\xed\x55\xda\xee\x2c\x19\xd7\x3f\xe2\xd8\xb5\xad\xf8\x62\xcd\xf7\x74\xf7\x01\xc7\x2e\xca\xf0\x9c\xe4\x1d\x3a\x3e\x7f\xb8\x31\x7b\xa3\x83\x11\x06\x57\x38\x00\x96\x96\xf4\x34\x79\xb0\xc1\xf8\x50\x48\x24\xf2\xf9\xad\xb7\x7b\x93\xaa\xbd\xf2\xd4\xe1\xaf\xfe\xb3\x7b\x17\xab\x55\x33\x45\x13\x38\xcf\x1a\x65\x14\xee\xa8\x66\x42\x39\x08\x2a\x05\x19\x99\x01\x4d\x68\x79\x5b\x08\x18\x83\x4e\x4e\x48\x1b\xbe\x29\xee\x7c\x92\x1e\x4a\x6f\x0e\xd9\xdf\x46\x5a\xe8\xed\xe0\xbe\xef\x8e\xd8\x3a\xe3\xbe\x2b\x19\x4f\xb2\x36\xce\x78\xc2\x09\x6e\x13\x1d\x5a\xbf\xeb\xe4\x7f\xf5\x1c\x8b\x96\x04\xde\x7a\x01\xac\xd9\x35\xe5\x04\xc8\x64\x38\x06\x4b\xda\x07\x77\x3b\x66\x06\xdf\xbb\x77\x77\xf8\x98\x8f\x64\xb5\xe3\x6f\x38\xfd\xfb\xed\x98\x16\xbc\xfe\x1f\x12\x52\x5e\xf9\xc9\x26\x53\x7b\x8d\xe6\x42\xf2\xf9\xe2\xe9\xff\x04\x00\x00\xff\xff\xf4\xd0\x74\x5b\x30\x4f\x00\x00")

func templatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_templatesJson,
		"templates.json",
	)
}

func templatesJson() (*asset, error) {
	bytes, err := templatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.json", size: 20272, mode: os.FileMode(420), modTime: time.Unix(1607715625, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates.json": templatesJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates.json": &bintree{templatesJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

