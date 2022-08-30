#!/bin/bash



function publish_rpc(){
  local TRG_FILE=${1:?"Target file required"}
  dasel -p toml put string  -f "${TRG_FILE}" -s ".rpc.laddr" -v "tcp://0.0.0.0:26657"

  # dasel seem to not support setting lists with toml
  cp "${TRG_FILE}" "${TRG_FILE}.tmp"
  dasel -r toml -w json -f ${TRG_FILE}.tmp | \
    jq 'setpath(path(.rpc.cors_allowed_origins);["*"])'  | \
    dasel -r json -w toml > ${TRG_FILE}
  rm ${TRG_FILE}.tmp

}
function update_persistent_peers() {
  local SRC_FILE=${1:?"Source file required"}
  local TRG_FILE=${2:?"Target file required"}

  local PEERS=$(dasel -p toml -f "${SRC_FILE}" -s ".p2p.persistent_peers")
  dasel -p toml put string  -f "${TRG_FILE}" -s ".p2p.persistent_peers" -v $PEERS
}

function update_config_toml(){
  local SRC=${1:?"Source folder required"}
  local TRG=${2:?"Target folder required"}
  # dasel strips comments for now. https://github.com/TomWright/dasel/issues/178
  # creating a backup
  local CONFIG_FILE_NAME="config/config.toml"
  local TRG_FILE=${TRG%/}/$CONFIG_FILE_NAME
  local SRC_FILE=${SRC%/}/$CONFIG_FILE_NAME
  cp "${TRG_FILE}" "${TRG_FILE}.bak"
  echo "cp "${TRG_FILE}" "${TRG_FILE}.bak""
  update_persistent_peers "${SRC_FILE}" "${TRG_FILE}"
  publish_rpc "${TRG_FILE}"
}