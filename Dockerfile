FROM alpine:latest
LABEL maintainer="chris@cbrgm.de"
LABEL version="latest"

ARG HUGO_VERSION

COPY ./drone-hugo.sh /bin/

RUN apk update
RUN chmod +x bin/drone-hugo.sh
RUN mkdir /temp/
RUN wget -O- https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz | tar xz -C /temp/
RUN mv /temp/hugo /bin/hugo
RUN rm  -rf /temp

ENTRYPOINT /bin/sh /bin/drone-hugo.sh
