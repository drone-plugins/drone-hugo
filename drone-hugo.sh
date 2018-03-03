#!/bin/bash
# Author: Christian Bargmann <chris@cbrgm.de>
set -e

# Assign default variables if not set
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

function addArgument {
  echo $1
  HUGO_COMMAND="${HUGO_COMMAND} $1"
}

# Hugo Command
HUGO_COMMAND="hugo"

# Validate config file
if [[ $PLUGIN_VALIDATE = true ]]; then
  echo "Checking config file ${PLUGIN_CONFIG}..."
  if [[ $PLUGIN_CONFIG != "false" ]]; then
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
