#!/bin/bash
SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/toml-helpers.sh

CONFIG_FILE_NAME="config/app.toml"

function enable_rest() {
  TRG=${1:?"Target folder required"}

  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "api" "enable" "true"
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "api" "swagger" "true"
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "api" "enabled-unsafe-cors" "true"
}

function update_app_toml(){
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}

  enable_rest "$TRG"
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "grpc-web" "enabled-unsafe-cors" "true"
}