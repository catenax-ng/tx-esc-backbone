#!/bin/bash

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/template-helpers.sh
source $SCRIPT_LOCATION/service-template.sh
source $SCRIPT_LOCATION/networks-template.sh
source $SCRIPT_LOCATION/volume-template.sh

NODE_SERVICE_PLACEHOLDER={esc-backbone-node}
VOLUME_PREFIX="chain"
NETWORK_NAME="validator"
IP_VERSION="4"

function apply_on_each (){
  array=($@)
  func=$1
  params=("${array[@]:1}")
  for entry in "${params[@]}"
  do
    $func "$entry"
  done
}

function add_node_service(){
  local HOME_DIR="${1:?"Home folder required"}"
  node_service $HOME_DIR $VOLUME_PREFIX $NETWORK_NAME $IP_VERSION
}

function add_node_services() {
  apply_on_each add_node_service $@
}

function add_node_volume(){
  local HOME_DIR="${1:?"Home folder required"}"
  node_volume $HOME_DIR $DOCKER_COMPOSE_FILE $VOLUME_PREFIX
}

function add_node_volumes() {
  local HOME_DIRS=($@)
  echo "volumes:\
  "
  apply_on_each add_node_volume ${HOME_DIRS[@]}
}

function generate_docker_compose_yml () {
  echo "$@"
  local DOCKER_COMPOSE_FILE=${1:?"Path to output docker-compose.yml required"}
  local NODES_FOLDER=${2:?"Folder containing the nodes home folders required"}
  local NETWORK_ADDRESS="${3:?"Network address required"}"
  local GATEWAY_ADDRESS="${4:?"Gateway address required"}"
  NODE_HOMES=(${NODES_FOLDER%/}/*/)
  cat "${SCRIPT_LOCATION/}/docker-compose-begin" > $DOCKER_COMPOSE_FILE
  add_node_services ${NODE_HOMES[@]} >> "$DOCKER_COMPOSE_FILE"
  add_networks_node $NETWORK_NAME "$NETWORK_ADDRESS" "$GATEWAY_ADDRESS" >> "$DOCKER_COMPOSE_FILE"
  add_node_volumes ${NODE_HOMES[@]}  >> "$DOCKER_COMPOSE_FILE"
}