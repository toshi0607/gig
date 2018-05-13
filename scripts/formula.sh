#!/bin/sh
set -eu
ROOT_DIR=${SCRIPTS}/..
ARTIFACTS_DIR=${ROOT_DIR}/dist/snapshot
PACKAGE=gig
PACKAGE_FULL=github.com/toshi0607/${PACKAGE}
FORMULA_CLASS=Gig
VERSION=$(gobump show -r)
TAG=v${VERSION}

shasum256() {
  local os=${1}
  local arch=${2}

  shasum -a 256 ${ARTIFACTS_DIR}/${PACKAGE}_${os}_${arch}.zip | awk '{print $1}'
}

formula(){
  cat << EOF > ${PACKAGE}_formula.rb
require "formula"
class ${FORMULA_CLASS} < Formula
  homepage 'https://${PACKAGE_FULL}'
  version '${VERSION}'
  if Hardware::CPU.is_32_bit?
    if OS.linux?
      url 'https://${PACKAGE_FULL}/releases/download/${TAG}/${PACKAGE}_linux_386.zip'
      sha256 '$(shasum256 linux 386)'
    else
      url 'https://${PACKAGE_FULL}/releases/download/${TAG}/${PACKAGE}_darwin_386.zip'
      sha256 '$(shasum256 darwin 386)'
    end
  else
    if OS.linux?
      url 'https://${PACKAGE_FULL}/releases/download/${TAG}/${PACKAGE}_linux_amd64.zip'
      sha256 '$(shasum256 linux amd64)'
    else
      url 'https://${PACKAGE_FULL}/releases/download/${TAG}/${PACKAGE}_darwin_amd64.zip'
      sha256 '$(shasum256 darwin amd64)'
    end
  end
  def install
    bin.install '${PACKAGE}'
  end
end
EOF
}

formula