# drone-hugo

Automatically create static web page files using [hugo](https://github.com/gohugoio/hugo) within your drone pipeline! For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-hugo/).

## Build

Build the binaries with the following commands:

```go
make build
```

## Docker

Build the Docker image with the following commands:

```bash
make [amd64,arm64,amd] hugo=0.00.0
```

### Usage

```bash
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
  -e PLUGIN_VALIDATE=false \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/hugo:latest
```

