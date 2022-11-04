#!/bin/bash


function update_client_toml(){
  local SRC=${1:?"Source folder required"}
  local TRG=${2:?"Target folder required"}
  # dasel strips comments for now. https://github.com/TomWright/dasel/issues/178
  # creating a backup
  local CONFIG_FILE_NAME="config/client.toml"
  local TRG_FILE="${TRG%/}/$CONFIG_FILE_NAME"
  local SRC_FILE="${SRC%/}/$CONFIG_FILE_NAME"
  cp "${TRG_FILE}" "${TRG_FILE}.bak"
  echo "cp "${TRG_FILE}" "${TRG_FILE}.bak""
}