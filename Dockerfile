# Build Container
FROM golang:latest AS build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ADD . /go/src/github.com/mayflower/docker-ls
WORKDIR /go/src/github.com/mayflower/docker-ls
RUN set -ex \
  && go generate  github.com/mayflower/docker-ls/lib/... \
  && go build     github.com/mayflower/docker-ls/cli/docker-ls \
  && go build     github.com/mayflower/docker-ls/cli/docker-rm

# Target container that is produced by docker build
FROM alpine:latest
RUN set -ex \
  && apk add --no-cache ca-certificates
LABEL MAINTAINER="Mayflower GmbH"
COPY --from=build /go/src/github.com/mayflower/docker-ls/docker-* /bin/
