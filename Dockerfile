FROM alpine:latest
LABEL maintainer="chris@cbrgm.de"
LABEL version="latest"

ENV HUGO_VERSION="0.37"

COPY ./drone-hugo.sh /bin/
RUN chmod +x /bin/drone-hugo.sh

RUN apk update
RUN mkdir /temp/
RUN wget https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz -P /temp
RUN tar xzvf /temp/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz -C /temp/
RUN mv /temp/hugo /bin/hugo
RUN rm  -rf /temp

ENTRYPOINT /bin/sh /bin/drone-hugo.sh
