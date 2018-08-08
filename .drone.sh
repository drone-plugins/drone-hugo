#!/bin/sh

set -e
set -x

# compile the main binary
GOOS=linux GOARCH=amd64 CGO_ENABLED=0         go build -ldflags "-X main.build=${DRONE_BUILD_NUMBER}" -a -tags netgo -o release/linux/amd64/drone-hugo github.com/drone-plugins/drone-hugo/cmd/drone-hugo
GOOS=linux GOARCH=arm64 CGO_ENABLED=0         go build -ldflags "-X main.build=${DRONE_BUILD_NUMBER}" -a -tags netgo -o release/linux/arm64/drone-hugo github.com/drone-plugins/drone-hugo/cmd/drone-hugo
GOOS=linux GOARCH=arm   CGO_ENABLED=0 GOARM=7 go build -ldflags "-X main.build=${DRONE_BUILD_NUMBER}" -a -tags netgo -o release/linux/arm/drone-hugo   github.com/drone-plugins/drone-hugo/cmd/drone-hugo
GOOS=linux GOARCH=386   CGO_ENABLED=0         go build -ldflags "-X main.build=${DRONE_BUILD_NUMBER}" -a -tags netgo -o release/linux/i386/drone-hugo  github.com/drone-plugins/drone-hugo/cmd/drone-hugo