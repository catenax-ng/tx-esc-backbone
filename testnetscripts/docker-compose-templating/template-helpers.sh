#!/bin/bash

function template_for_script() {
    local SCRIPT_NAME=$( basename ${1:?"Path of the script required."} )
    cat "${SCRIPT_LOCATION%/}/${SCRIPT_NAME%.*}"
}

function index_for_node() {
  local HOME_DIR=${1:?"Home folder required"}
  cat "${HOME_DIR%/}/node_index"
}

function ip_address_for_node() {
  local HOME_DIR=${1:?"Home folder required"}
  cat "${HOME_DIR%/}/ip_address"
}