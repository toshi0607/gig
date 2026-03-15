#!/bin/bash
set -eu

go test -race -coverprofile=coverage.txt -covermode=atomic ./...
