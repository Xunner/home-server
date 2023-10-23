#!/usr/bin/env bash

export GO111MODULE=on

RUN_NAME="home-server"
mkdir -p output/
cp -r static/ output/static
go build -mod=mod -gcflags "all=-N -l" -o output/${RUN_NAME}
