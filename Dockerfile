FROM golang:latest
MAINTAINER Mayflower GmbH

ADD . /go/src/github.com/mayflower/docker-ls
WORKDIR /go/src/github.com/mayflower/docker-ls

RUN go generate github.com/mayflower/docker-ls/lib/... && go install github.com/mayflower/docker-ls/cli/...
