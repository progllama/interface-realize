FROM golang:alpine
RUN apk update && apk add git make gcc openssh