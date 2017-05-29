#!/usr/bin/env bash

export GIN_MODE=release && env GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o eta-amd -v github.com/qbunt/eta-announce-go
export GIN_MODE=release && env GOOS=linux GOARCH=arm GIN_MODE=release go build -o eta-arm -v github.com/qbunt/eta-announce-go
env GIN_MODE=release GOOS=linux GOARCH=arm GOARM=7 go build -o eta-armv7 -v github.com/qbunt/eta-announce-go
env GIN_MODE=release GOOS=linux GOARCH=arm GOARM=6 go build -o eta-armv6 -v github.com/qbunt/eta-announce-go
export GIN_MODE=release && env GIN_MODE=release go build -o eta-mac -v github.com/qbunt/eta-announce-go