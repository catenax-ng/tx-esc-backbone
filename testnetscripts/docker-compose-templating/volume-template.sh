#!/bin/bash

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/template-helpers.sh

function node_volume(){
  local HOME_DIR="${1:?"Home folder required"}"
  local DOCKER_COMPOSE_FILE="${2:?"Location of the docker-compose.yml required"}"
  local VOLUME_PREFIX="${3:-chain}"
  local NODE_INDEX=$(index_for_node "$HOME_DIR")

  local VOLUME_LOCATION="$(realpath --relative-to="$(dirname $DOCKER_COMPOSE_FILE)" "$HOME_DIR")"
  echo "$(template_for_script "${BASH_SOURCE[0]}")" | \
    sed 's/{node-index}/'"$NODE_INDEX"'/' | \
    sed "s/{volume-prefix}/$VOLUME_PREFIX/" | \
    sed "s/{volume-location}/${VOLUME_LOCATION//\//\\/}/"
}
