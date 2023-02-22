#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0


SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )

id
source $SCRIPT_LOCATION/cosmos-helpers.sh
source $SCRIPT_LOCATION/init-global-vars.sh
if [ -z "$GIT_REPO" ]; then
  echo "\$GIT_REPO has to be set"
  exit 1
fi

ORCHESTRATOR_HOME="${1:?Provide a home folder}"
CREATE_LOCAL_REPO="${2:-FALSE}"

if [ "$CREATE_LOCAL_REPO" = "TRUE" ]; then
  url="://"
  if [ -z "${CREATE_LOCAL_REPO##*$url*}" ] ;then
    echo "Cannot create a temporary repository for $CREATE_LOCAL_REPO"
  else
    if [ -d $GIT_REPO ]; then
      cd $GIT_REPO
      if [ ! $(git rev-parse --is-inside-work-tree 2> /dev/null) ]; then
        cd - > /dev/null
        create_a_local_empty_repo $GIT_REPO $REPO_BRANCH
      else
        cd - > /dev/null
      fi
    else
      create_a_local_empty_repo $GIT_REPO $REPO_BRANCH
    fi
  fi
fi

echo "Operating on $REPO_BRANCH at $GIT_REPO"
function publish_initial_genesis_file() {
  echo "publish_initial_genesis_file at $ORCHESTRATOR_HOME"
  clone_sync_repo "$ORCHESTRATOR_HOME" "$GIT_REPO"
  sync_generate_genesis_file "$ORCHESTRATOR_HOME"
}

function add_faucet_account(){
  echo "add_faucet_account"
  if  [ "i-know-this-is-insecure" == "$ADD_FAUCET_ACCOUNT" ]; then
    echo "!!!Adding a faucet for mnemonic: $FAUCET_MNEMONIC"
    local FAUCET_DIR="${ORCHESTRATOR_HOME%/}/faucet"
    mkdir $FAUCET_DIR
    local MONIKER=$(home_name "$FAUCET_DIR")
    FAUCET_MNEMONIC_FILE=$FAUCET_DIR/mnemonic
    echo "$FAUCET_MNEMONIC" > $FAUCET_MNEMONIC_FILE
    each_add_key $FAUCET_DIR $MONIKER "$FAUCET_MNEMONIC"
    cd "${ORCHESTRATOR_HOME}"
    git add faucet
    git commit -m "Add faucet account"
    cd - > /dev/null
    local PUB_ADDR=$(get_pub_addr $FAUCET_DIR $MONIKER)
    sed -i "s/\s\{0,\}\#\{0,\}\s\{0,\}moniker\s\{0,\}=.\{0,\}/moniker = faucet/" "${FAUCET_DIR%/}/config/config.toml"
    add_funds_to_addr "$ORCHESTRATOR_HOME" $PUB_ADDR $FAUCET_INITIAL_BALANCE $CURRENCY
  fi
}

function add_funds_for_initial_account() {
  echo "add_funds_for_initial_account for $@"
  local PUB_ADDR_FILE=${1:?"Requires a file with the public address as name and the moniker as content."}
  local PUB_ADDR=$(basename $PUB_ADDR_FILE)
  local MONIKER=$(cat $PUB_ADDR_FILE)
  add_funds_to_addr "$ORCHESTRATOR_HOME" $PUB_ADDR $VALIDATOR_INITIAL_BALANCE $CURRENCY
}

function add_initial_accounts() {
  echo "add_initial_accounts at $ORCHESTRATOR_HOME"
  local PUB_ADDR_GLOB="$ORCHESTRATOR_HOME/pub_addr/*"
  wait_for_validator_commits $ORCHESTRATOR_HOME "$PUB_ADDR_GLOB" $VALIDATOR_COUNT
  local PUB_ADDRS=($PUB_ADDR_GLOB)
  apply_on_each add_funds_for_initial_account ${PUB_ADDRS[@]}
  add_faucet_account
  tag_genesis_file "$ORCHESTRATOR_HOME" "genesis_with_funds" "add funds to genesis file"
}

function apply_gentxs(){
    local GENTX_GLOB="$ORCHESTRATOR_HOME/config/gentx/gentx-*"
    wait_for_validator_commits $ORCHESTRATOR_HOME "$GENTX_GLOB" $VALIDATOR_COUNT
    echo "$CHAIN_BINARY --home $ORCHESTRATOR_HOME collect-gentxs"
    $CHAIN_BINARY --home "$ORCHESTRATOR_HOME" collect-gentxs
    tag_genesis_file $ORCHESTRATOR_HOME "genesis_with_txs" "genesis file with transactions"
}


publish_initial_genesis_file
add_initial_accounts
apply_gentxs