FROM golang:alpine
RUN apk update && apk add make
COPY . /go/src/interface-realize