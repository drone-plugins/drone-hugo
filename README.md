# drone-hugo

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-hugo/status.svg)](http://cloud.drone.io/drone-plugins/drone-hugo)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/hugo.svg)](https://microbadger.com/images/plugins/hugo "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-hugo?status.svg)](http://godoc.org/github.com/drone-plugins/drone-hugo)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-hugo)](https://goreportcard.com/report/github.com/drone-plugins/drone-hugo)

Automatically create static web page files using [Hugo](https://github.com/gohugoio/hugo) within your Drone pipeline. For the usage information and a listing of the available options please take a look at [the docs](https://plugins.drone.io/plugins/hugo).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-hugo
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/hugo .
```

### Usage

```console
docker run --rm \
  -e PLUGIN_HUGO_VERSION=0.00.0 \
  -e PLUGIN_BUILDDRAFTS=false \
  -e PLUGIN_BUILDEXPIRED=false \
  -e PLUGIN_BUILDFUTURE=false \
  -e PLUGIN_CACHEDIR=false \
  -e PLUGIN_CONFIG=false \
  -e PLUGIN_CONTENT=false \
  -e PLUGIN_LAYOUT=false \
  -e PLUGIN_OUTPUT=false \
  -e PLUGIN_SOURCE=false \
  -e PLUGIN_THEME=false \
  -e PLUGIN_OUTPUT=false \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/hugo:latest
```

