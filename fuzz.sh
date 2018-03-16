#!/bin/sh

set -e -x

mkdir -p corpus
go test -short ./...
go-fuzz-build github.com/buildkite/shellwords
go-fuzz -bin=./shellwords-fuzz.zip -workdir=.
