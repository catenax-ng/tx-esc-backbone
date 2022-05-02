#!/bin/bash

CONFIG_FILE_NAME="config/config.toml"
function update_persistent_peers() {
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}

  REPLACEMENT="$(cat ${SRC%/}/$CONFIG_FILE_NAME | grep "persistent_peers = \"")"
  REPLACEMENT=${REPLACEMENT//\//\\/}
  REPLACEMENT=${REPLACEMENT//:/\\:}
  sed -i "s/\s\{0,\}\#\{0,\}\s\{0,\}persistent_peers\s\{0,\}=.*/$REPLACEMENT/" "${TRG%/}/$CONFIG_FILE_NAME"
}


function update_config_toml(){
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}
  update_persistent_peers "$SRC" "$TRG"
}