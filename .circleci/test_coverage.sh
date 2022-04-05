#!/usr/bin/env sh

# run cover analysis for whole project
set -e
mkdir -p /build/fsm/go_test/
echo "" > /build/fsm/go_test/coverage.txt
export CGO_ENABLED=0
echo "running go test coverage"
go test -covermode=count -coverpkg=./... -coverprofile=/build/fsm/go_test/coverage.txt ./...
echo "running go test coverage complete"
