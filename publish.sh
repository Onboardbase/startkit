#!/bin/bash

set -e

VERSION="$1"

if [[ -z "$VERSION" ]]; then
  echo "Need a version"
  exit 1
fi

GOOS_LIST="linux darwin"
GOARCH_LIST="amd64 arm64"

rm -rf npm/bin/*
for goos in $GOOS_LIST; do
  for goarch in $GOARCH_LIST; do
    CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" go build -ldflags="-s -w -X main.Version=$VERSION" -o "npm/bin/gen-gopm3-${goos}-${goarch}"

    # There appears to be issues with macOS and upx compressed binaries.
    if [[ "$goos" != "darwin" ]]; then
      upx "npm/bin/gen-gopm3-${goos}-${goarch}"
    fi
  done
done

cat << EOF > npm/package.json
{
  "name": "onboardbasekit",
  "version": "$VERSION",
  "description": "**Onboardbase Kit** (onboardbasekit) is a CLI tool for generating boilerplates with onboardbase preconfigured in them as a SecretOps infrastucture.",
  "main": "src/index.js",
  "scripts": {
    "postinstall": "node install.js"
  },
  "bin": {
    "gopm3": "src/index.js"
  },
  "repository": {
    "type": "git",
    "url": "git+https://github.com/Onboardbase/onboardbasekit.git"
  },
  "author": "Samuel Olusola",
  "license": "MIT",
  "bugs": {
    "url": "https://github.com/Onboardbase/onboardbasekit/issues"
  },
  "homepage": "https://github.com/Onboardbase/onboardbasekit#readme"
}
EOF

git add npm/package.json
git commit -m "Publish version ${VERSION}"
git tag "v${VERSION}" -f
git push
git push --tags
pushd npm
npm publish --access public .
popd