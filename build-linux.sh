#!/usr/bin/env bash

GOPROXY=direct GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"