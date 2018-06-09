#!/bin/bash
set -eu

curl -X POST -d '{"owner": "toshi0607", "repo": "gig"}' ${RELEASE_TWEETER_ENDPOINT}
