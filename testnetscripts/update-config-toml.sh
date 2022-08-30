#!/bin/bash
SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/toml-helpers.sh

CONFIG_FILE_NAME="config/config.toml"

function publish_rpc(){
  TRG=${1:?"Target folder required"}
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "rpc" "laddr" "\"tcp://127.0.0.1:26657\""
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "rpc" "cors_allowed_origins" "[\"*\"]"
}
function update_persistent_peers() {
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}

  REPLACEMENT="$(cat ${SRC%/}/$CONFIG_FILE_NAME | grep "persistent_peers = \"")"
  REPLACEMENT=${REPLACEMENT//\//\\/}
  REPLACEMENT=${REPLACEMENT//:/\\:}
  sed -i "s/\s\{0,\}\#\{0,\}\s\{0,\}persistent_peers\s\{0,\}=.\{0,\}/$REPLACEMENT/" "${TRG%/}/$CONFIG_FILE_NAME"
}

# cors_allowed_origins = ["*"]
function update_config_toml(){
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}
  update_persistent_peers "$SRC" "$TRG"
}