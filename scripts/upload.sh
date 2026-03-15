#!/bin/sh
set -eu

# set GITHUB_TOKEN=...

goreleaser release --clean
