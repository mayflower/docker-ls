FROM golang:latest
MAINTAINER Mayflower GmbH

ADD . /go/src/github.com/mayflower/docker-ls
WORKDIR /go/src/github.com/mayflower/docker-ls

RUN make clean && make install && cp /go/src/github.com/mayflower/docker-ls/build/bin/* /usr/local/bin/
