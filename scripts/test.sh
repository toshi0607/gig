#!/bin/bash
set -eu

SCRIPTS=$(cd $(dirname $0) && pwd)

echo "" > coverage.txt
for d in $(${SCRIPTS}/packages.sh); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
