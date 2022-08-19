### Webapp build
FROM node:16.3.0@sha256:ca6daf1543242acb0ca59ff425509eab7defb9452f6ae07c156893db06c7a9a4 as nodebuild

WORKDIR /webapp
ADD webapp/ /webapp

RUN npm install --no-optional && \
    npm run pack

### mvsqlite
FROM curlimages/curl:latest as mvsqlite
RUN curl -L -o /tmp/libmvsqlite_preload.so https://github.com/losfair/mvsqlite/releases/download/v0.1.13/libmvsqlite_preload.so

### Go build
FROM golang:1.18.5-bullseye AS gobuild

WORKDIR /go/src/focalboard
ADD . /go/src/focalboard


RUN apt update && apt install -y libsqlite3-dev
RUN  EXCLUDE_PLUGIN=true EXCLUDE_SERVER=true EXCLUDE_ENTERPRISE=true make server-linux
RUN mkdir /data

## Final image
FROM debian:bullseye-slim

WORKDIR /opt/focalboard

COPY --from=gobuild --chown=nobody:nogroup /data /data
COPY --from=nodebuild --chown=nobody:nogroup /webapp/pack pack/
COPY --from=gobuild --chown=nobody:nogroup /go/src/focalboard/bin/linux/focalboard-server bin/
COPY --from=gobuild --chown=nobody:nogroup /go/src/focalboard/LICENSE.txt LICENSE.txt
COPY --from=gobuild --chown=nobody:nogroup /go/src/focalboard/docker/server_config.json config.json
COPY --from=mvsqlite /tmp/libmvsqlite_preload.so /
RUN apt update && apt install -y libsqlite3-0 libssl1.1 ca-certificates

USER nobody

EXPOSE 8000/tcp

EXPOSE 8000/tcp 9092/tcp

VOLUME /data

ENV LD_PRELOAD=/libmvsqlite_preload.so
CMD ["/opt/focalboard/bin/focalboard-server"]
