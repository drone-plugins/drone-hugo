#!/bin/bash
# Author: Christian Bargmann <chris@cbrgm.de>
set -e

# Assign default variables if not set
PLUGIN_VERSION=${PLUGIN_VERSION:-"false"}
PLUGIN_BUILDDRAFTS=${PLUGIN_BUILDDRAFTS:-"false"}
PLUGIN_BUILDEXPIRED=${PLUGIN_BUILDEXPIRED:-"false"}
PLUGIN_BUILDFUTURE=${PLUGIN_BUILDFUTURE:-"false"}
PLUGIN_CONFIG=${PLUGIN_CONFIG:-"false"}
PLUGIN_CONTENT=${PLUGIN_CONTENT:-"false"}
PLUGIN_LAYOUT=${PLUGIN_LAYOUT:-"false"}
PLUGIN_OUTPUT=${PLUGIN_OUTPUT:-"false"}
PLUGIN_SOURCE=${PLUGIN_SOURCE:-"false"}
PLUGIN_THEME=${PLUGIN_THEME:-"false"}
PLUGIN_URL=${PLUGIN_URL:-"false"}
PLUGIN_VALIDATE=${PLUGIN_VALIDATE:-"false"}

# The hugo command
HUGO_COMMAND="hugo"

function addArgument {
  echo $1
  HUGO_COMMAND="${HUGO_COMMAND} $1"
}

# Download hugo binary if version attribute is set. When not set, use the one shipped binary inside container
if [[ $PLUGIN_VERSION != "false" ]] ; then
  echo "Downloading hugo version v${PLUGIN_VERSION}..."
  mkdir /temp/
  wget -qO- https://github.com/gohugoio/hugo/releases/download/v${PLUGIN_VERSION}/hugo_${PLUGIN_VERSION}_Linux-64bit.tar.gz | tar xz -C /temp/
  mv /temp/hugo /bin/hugo
  rm  -rf /temp
  echo "Using hugo v${PLUGIN_VERSION}..."
fi

# Validate config file
if [[ $PLUGIN_VALIDATE = true ]]; then
  if [[ $PLUGIN_CONFIG != "false" ]]; then
    echo "Checking config file ${PLUGIN_CONFIG}..."
    hugo check --config ${PLUGIN_CONFIG}
  else
    hugo check
  fi
fi

# Create hugo command from arguments
if [[ $PLUGIN_BUILDDRAFTS != "false" ]] ; then addArgument "-D" ; fi
if [[ $PLUGIN_BUILDEXPIRED != "false" ]] ; then addArgument "-E" ; fi
if [[ $PLUGIN_BUILDFUTURE != "false" ]] ; then addArgument "-F" ; fi
if [[ $PLUGIN_CONFIG != "false" ]] ; then addArgument "--config ${PLUGIN_CONFIG}" ; fi
if [[ $PLUGIN_CONTENT != "false" ]] ; then addArgument "--contentDir ${PLUGIN_CONFIG}" ; fi
if [[ $PLUGIN_LAYOUT != "false" ]] ; then addArgument "--layoutDir ${PLUGIN_LAYOUT}" ; fi
if [[ $PLUGIN_OUTPUT != "false" ]] ; then addArgument "--destination ${PLUGIN_OUTPUT}" ; fi
if [[ $PLUGIN_SOURCE != "false" ]] ; then addArgument "--source ${PLUGIN_SOURCE}" ; fi
if [[ $PLUGIN_THEME != "false" ]] ; then addArgument "--theme ${PLUGIN_THEME}" ; fi
if [[ $PLUGIN_URL != "false" ]] ; then addArgument "--baseURL ${PLUGIN_URL}" ; fi

echo "Executing: ${HUGO_COMMAND}"
$HUGO_COMMAND
