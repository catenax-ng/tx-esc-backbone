#!/bin/bash
SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/toml-helpers.sh

CONFIG_FILE_NAME="config/client.toml"
function update_client_toml(){
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}

}