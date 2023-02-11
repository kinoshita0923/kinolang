FROM golang:1.18.10-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir -p /src/kinolang

COPY . /src/kinolang

COPY ~/.ssh ~/.ssh

WORKDIR /src/kinolang