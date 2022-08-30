#!/bin/bash

local CONFIG_FILE_NAME="config/app.toml"

function enable_rest() {
  local TRG_FILE=${1:?"Target file required"}
  dasel -p toml put bool  -f "${TRG_FILE}" -s ".api.enable" -v "true"
  dasel -p toml put bool  -f "${TRG_FILE}" -s ".api.swagger" -v "true"
  dasel -p toml put bool  -f "${TRG_FILE}" -s ".api.enabled-unsafe-cors" -v "true"
}

function update_app_toml(){
  local SRC=${1:?"Source folder required"}
  local TRG=${2:?"Target folder required"}
  # dasel strips comments for now. https://github.com/TomWright/dasel/issues/178
  # creating a backup
  local CONFIG_FILE_NAME="config/app.toml"
  local TRG_FILE="${TRG%/}/config/app.toml"
  local SRC_FILE="${SRC%/}/config/app.toml"
  cp "${TRG_FILE}" "${TRG_FILE}.bak"
  echo "cp "${TRG_FILE}" "${TRG_FILE}.bak""
  enable_rest "${TRG_FILE}"
  dasel -p toml put bool  -f "${TRG_FILE}" -s ".grpc-web.enable-unsafe-cors" -v "true"
}