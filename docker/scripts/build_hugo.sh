#! /bin/sh

set -e

hugo_version=${HUGO_VERSION:-0.87.0}

echo "downloading build dependencies..."
apk add -U --no-cache git build-base

LDFLAGS="-s -w -X github.com/gohugoio/hugo/common/hugo.buildDate=\$(date +%Y-%m-%dT%H:%M:%SZ) -X github.com/gohugoio/hugo/common/hugo.commitHash=\$(git rev-parse --short HEAD)"

echo "downloading hugo version v$hugo_version..."
git clone --branch "v$hugo_version" https://github.com/gohugoio/hugo.git

echo "building hugo..."
cd hugo

CGO_ENABLED=0 go build -ldflags "\"$LDFLAGS\"" -o /tmp/hugo .

CGO_ENABLED=1 go build -tags extended -ldflags "\"$LDFLAGS\"" -o /tmp/hugo-extended

echo "done."
