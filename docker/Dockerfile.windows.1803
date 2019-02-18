# escape=`
FROM plugins/base:windows-1803

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" `
  org.label-schema.name="Drone Mercurial" `
  org.label-schema.vendor="Drone.IO Community" `
  org.label-schema.schema-version="1.0"

# TODO: install required tools

ADD release\drone-hugo.exe c:\drone-hugo.exe
ENTRYPOINT [ "c:\\drone-hugo.exe" ]
