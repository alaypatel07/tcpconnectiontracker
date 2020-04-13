FROM golang:1.13

WORKDIR /go/src/github.com/alaypatel07/tcpconnectiontracker
COPY . .
ENV GO_PACKAGE github.com/alaypatel07/tcpconnectiontracker
USER 0
RUN go build
