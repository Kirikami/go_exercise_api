FROM golang:latest

MAINTAINER Ruslan Naumenko

## Copy project folder
COPY . /go/src/github.com/kirikami/go_exercise_api/

## Set working directory
WORKDIR /go/src/github.com/kirikami/go_exercise_api

RUN ln -s /go/src/github.com/kirikami/go_exercise_api/run_app.sh /usr/bin/run_app

## Install glide and dependency
COPY images/go/glide /bin/
RUN glide install

## Install goose
RUN go get bitbucket.org/liamstask/goose/cmd/goose

## Install application
RUN go install