#!/usr/bin/env sh

set -euf

echo --- :go: Checking go mod tidyness
go mod tidy
if ! git diff --no-ext-diff --exit-code; then
  echo ^^^ +++
  echo "The go.mod or go.sum files are out of sync with the source code"
  echo "Please run \`go mod tidy\` locally, and commit the result."

  exit 1
fi

echo --- :go: Checking go formatting
gofmt -w .
if ! git diff --no-ext-diff --exit-code; then
  echo ^^^ +++
  echo "Files have not been formatted with gofmt."
  echo "Fix this by running \`go fmt ./...\` locally, and committing the result."

  exit 1
fi

echo --- :go: Generating code
go generate ./...
if ! git diff --no-ext-diff --exit-code; then
  echo ^^^ +++
  echo :x: Generated code was not commited.
  echo "Run"
  echo "  go generate ./..."
  echo "and make a commit."

  exit 1
fi

echo --- :broom: Downloading golangci-lint...
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v2.1.6

echo --- :go: Running golangci-lint...
./bin/golangci-lint run

echo +++ Everything is clean and tidy! 🎉
