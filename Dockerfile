FROM golang:1.18.10-alpine

RUN apk update && apk add git

RUN mkdir -p /src/kinolang

COPY . /src/kinolang

WORKDIR /src/kinolang