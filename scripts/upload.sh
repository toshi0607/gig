#!/bin/sh
set -eu

# set GITHUB_TOKEN=...

latest_tag=$(git describe --abbrev=0 --tags)
goxc
ghr -u toshi0607 -r gig $latest_tag dist/snapshot/
