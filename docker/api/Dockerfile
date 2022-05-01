FROM golang:latest AS build

WORKDIR /go/src/api
ENV GO111MODULE=on
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download
COPY ./api .

RUN go build -o /go/bin/app && \
  go install github.com/cosmtrek/air@latest