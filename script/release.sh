#!/bin/sh

# Usage: GITHUB_TOKEN=... script/release
set -e
latest_tag=$(git describe --abbrev=0 --tags)
goxc
ghr -u toshi0607 -r gig $latest_tag dist/snapshot/