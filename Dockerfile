FROM golang:latest

MAINTAINER Ruslan Naumenko

## Copy project folder
COPY . /go/src/github.com/kirikami/go_exercise_api/

## Set working directory
WORKDIR /go/src/github.com/kirikami/go_exercise_api

## Install glide and dependency
COPY images/go/glide /bin/
RUN glide install

## Install goose
RUN go get bitbucket.org/liamstask/goose/cmd/goose

## Install application
RUN go install
