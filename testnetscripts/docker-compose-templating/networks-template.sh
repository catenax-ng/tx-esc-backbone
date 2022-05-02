#!/bin/bash

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/template-helpers.sh

function add_networks_node(){
  local NETWORK_NAME="${1:?"Network name required"}"
  local NETWORK_ADDRESS="${2:?"Network address required"}"
  local GATEWAY_ADDRESS="${3:?"Gateway address required"}"
  echo "$(template_for_script ${BASH_SOURCE[0]})" | \
    sed "s/{network-name}/$NETWORK_NAME/" | \
    sed "s/{network-address}/${NETWORK_ADDRESS//\//\\/}/" | \
    sed "s/{gateway-address}/$GATEWAY_ADDRESS/"
}