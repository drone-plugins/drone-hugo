FROM plugins/base:linux-arm64

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Hugo" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

ENV HUGO_VERSION=0.46
ENV HUGO_ARCH=arm64
ENV PLUGIN_HUGO_ARCH=$HUGO_ARCH
ENV PLUGIN_HUGO_SHIPPED_VERSION=$HUGO_VERSION

RUN apk --no-cache add git && \
    wget -O- https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-${HUGO_ARCH}.tar.gz | tar xz -C /bin

ADD release/linux/arm64/drone-hugo /bin/
ENTRYPOINT ["/bin/drone-hugo"]
