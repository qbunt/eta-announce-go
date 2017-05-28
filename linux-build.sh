#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 export GIN_MODE=release go build -v github.com/qbunt/eta-announce-go -o eta-announce-go-amd
env GOOS=linux GOARCH=arm export GIN_MODE=release go build -o eta-announce-go-arm -v github.com/qbunt/eta-announce-go