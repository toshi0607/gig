#!/bin/sh
set -eu

ROOT_DIR=$(cd $(dirname $0) && cd .. && pwd)

echo current version: $(gobump show -r)
read -p "input next version: " next_version

gobump set $next_version -w
ghch -w -N v$next_version

${ROOT_DIR}/scripts/_scoop.sh

git add version.go CHANGELOG.md gig.json
git commit -m "Checking in changes prior to tagging of version v$next_version"
git tag v$next_version
git push && git push --tags
