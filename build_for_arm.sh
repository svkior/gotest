#!/bin/sh
GOARCH=arm GOARM=5 GOOS=linux go build src/gotest.go
GOARCH=arm GOARM=5 GOOS=linux go build src/test_mmap.go
