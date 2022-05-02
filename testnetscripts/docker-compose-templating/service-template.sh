#!/bin/bash

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/template-helpers.sh

API_PORT_STEP=1
API_BASE_PORT=1317
TENDERMINT_PORT_STEP=100
TENDERMINT_LOWER_BASE_PORT=26656
TENDERMINT_UPPER_BASE_PORT=26658


function node_service(){
  local HOME_DIR="${1:?"Home folder required"}"
  local VOLUME_PREFIX="${2:-chain}"
  local NETWORK_NAME="${3:-validator}"
  local IP_VERSION="${4:-4}"

  local NODE_INDEX=$(index_for_node "$HOME_DIR")
  local NETWORK_ADDRESS=$(ip_address_for_node "$HOME_DIR")

  local API_PORT=$(( ( $NODE_INDEX - 1 ) * $API_PORT_STEP + $API_BASE_PORT ))
  local TENDERMINT_LOWER_PORT=$(( ( $NODE_INDEX - 1 ) * $TENDERMINT_PORT_STEP + $TENDERMINT_LOWER_BASE_PORT ))
  local TENDERMINT_UPPER_PORT=$(( ( $NODE_INDEX - 1 ) * $TENDERMINT_PORT_STEP + $TENDERMINT_UPPER_BASE_PORT ))
  local TENDERMINT_PORTS="$TENDERMINT_LOWER_PORT-$TENDERMINT_UPPER_PORT"

  echo "$(template_for_script "${BASH_SOURCE[0]}")" | \
    sed 's/{node-index}/'"$NODE_INDEX"'/' | \
    sed "s/{volume-prefix}/$VOLUME_PREFIX/" | \
    sed "s/{network-name}/$NETWORK_NAME/" | \
    sed "s/{ipv}/$IP_VERSION/" | \
    sed "s/{address}/$NETWORK_ADDRESS/" | \
    sed "s/{api-port}/$API_PORT/" | \
    sed "s/{tendermint-ports}/$TENDERMINT_PORTS/"
}