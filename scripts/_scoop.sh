#!/bin/sh
set -eu

ROOT_DIR=$(cd $(dirname $0) && cd .. && pwd)
VERSION=$(gobump show -r)

cat << EOF > ${ROOT_DIR}/gig.json
{
  "version": "${VERSION}",
  "architecture": {
    "64bit": {
      "url": "https://github.com/toshi0607/gig/releases/download/v${VERSION}/gig_windows_amd64.zip",
      "bin": "gig.exe",
      "extract_dir": "gig_windows_amd64"
    },
    "32bit": {
      "url": "https://github.com/toshi0607/gig/releases/download/v${VERSION}/gig_windows_386.zip",
      "bin": "gig.exe",
      "extract_dir": "gig_windows_386"
    }
  },
  "homepage": "https://github.com/toshi0607/gig"
}
EOF
