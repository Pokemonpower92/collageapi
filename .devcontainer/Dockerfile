ARG GO_VERSION=1.21.3

FROM golang:${GO_VERSION}

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/collageapi
ENV HOST_NAME localhost
ENV PORT 8080

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 8080
