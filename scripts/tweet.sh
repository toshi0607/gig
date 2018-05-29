#!/bin/bash
set -eu

curl -X POST -d '{"owner": "toshi0607", "repo": "release-tweeter"}' ${RELEASE_TWEETER_ENDPOINT}
