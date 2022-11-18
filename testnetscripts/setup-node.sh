#!/bin/bash

# Copyright (c) 2022 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

echo "start setup_node.sh for $1"

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/cosmos-helpers.sh
source $SCRIPT_LOCATION/init-global-vars.sh

if [ -z "$GIT_REPO" ]; then
  echo "\$GIT_REPO has to be set"
fi
HOME_FOLDER="${1:?Provide a home folder}"
PUBLIC_HOST_NAME="${2:?Provide the validators public hostname/ ip address}"
MNEMONIC="${3:?Provide the validator\'s mnemonic}"

function each_init_home(){
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=${2:?"Moniker required"}
  $CHAIN_BINARY --home $HOME_DIR init $MONIKER &> /dev/null
}

function create_home_and_key() {
  echo "create_home_and_key for $1"
  local NODE_HOME="$1"
  local MONIKER="$(home_name "$NODE_HOME")"
  local REPO="${NODE_HOME%/}/sync"
  echo "Processing $MONIKER"
  clone_sync_repo "$REPO" "$GIT_REPO"
  echo "each_init_home "$NODE_HOME" "$MONIKER""
  each_init_home "$NODE_HOME" "$MONIKER"
  each_add_key "$NODE_HOME" "$MONIKER"
  echo rm "$NODE_HOME/config/genesis.json"
  rm "$NODE_HOME/config/genesis.json"
  each_fetch_genesis_file_for_node_from_tag "$NODE_HOME" "init_genesis"
}


function each_fetch_genesis_file_for_node_from_tag(){
  local HOME_DIR=${1:?"Home folder required"}
  local REPO="${HOME_DIR%/}/sync"
  local TAG=${2:?"tag name required"}
  wait_for_tag $HOME_DIR $TAG
  pull_git $REPO $TAG
  mkdir -p "${HOME_DIR%/}/config/"
  cp "$REPO/config/genesis.json" "${HOME_DIR%/}/config/"
}

function wait_for_tag(){
  local HOME_DIR=${1:?"Home folder required"}
  local REPO="${HOME_DIR%/}/sync"
  local TAG=${2:?"tag name required"}
  cd $REPO
  local retry=1
  git ls-remote $GIT_QUIET --exit-code --tags origin $TAG > /dev/null
  local lstag_result=$?
  while [ $lstag_result -ne 0 -a $retry -lt ${GIT_WAIT_MAX_RETRY:-5} ]
  do
    echo "no tag $TAG found"
    echo "waiting ${GIT_WAIT:-1}s ... "
    sleep ${GIT_WAIT:-1}
    retry=$(( $retry + 1 ))
    git ls-remote  $GIT_QUIET --exit-code --tags origin $TAG
    lstag_result=$?
  done
  if [ $lstag_result -ne 0 -a $retry -eq ${GIT_WAIT_MAX_RETRY:-5} ]
  then
    echo "no tag $TAG found"
    echo "retries exceeded"
    exit 1
  fi
  echo "tag $TAG found"
  cd - > /dev/null
}

function each_write_address_to_repo(){
  echo "each_write_address_to_repo for $1"
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=$(home_name "$HOME_DIR")
  local REPO="${HOME_DIR%/}/sync"
  echo "Publish public key for $MONIKER"
  pull_git $REPO
  mkdir -p $REPO/pub_addr
  local PUB_ADDR=$(get_pub_addr "$HOME_DIR" "$MONIKER")
  touch $REPO/pub_addr/$PUB_ADDR
  echo "$MONIKER" > $REPO/pub_addr/$PUB_ADDR
  cd $REPO
  git add pub_addr/$PUB_ADDR
  git commit $GIT_QUIET -s -m "Add public address for $MONIKER"
  push_git_retrying
  echo "Committed and pushed public address $REPO/pub_addr/$PUB_ADDR"
  cd - > /dev/null
}

function push_git_retrying() {
  if  [ ! $(git rev-parse --is-inside-work-tree) ]; then
    echo "Not in a git working folder."
    exit 1
  fi
  local BRANCH=${REPO_BRANCH:-main}
  echo "git push $(pwd)"
  git push $GIT_QUIET  2> /dev/null
  local push_result=$?
  local retry=1
  while [ $push_result -ne 0 -a $retry -lt ${GIT_PUSH_MAX_RETRY:-10} ]
  do
    retry=$(( $retry + 1 ))
    git fetch $GIT_QUIET origin
    git rebase $GIT_QUIET origin/$BRANCH
    git push $GIT_QUIET 2> /dev/null
    push_result=$?
  done
  if [ $push_result -ne 0 -a $retry -eq ${GIT_PUSH_MAX_RETRY:-10} ]
  then
      echo "Failed pushing the folder after $retry retries."
      echo "This might be caused, if the number of validators is greater than the value of \$GIT_PUSH_MAX_RETRY (default: \$VALIDATOR_COUNT + 2)"
      exit 2
  fi
}

function fetch_genesis_file_with_funds(){
  echo "fetch_genesis_file_with_funds for $1"
  each_fetch_genesis_file_for_node_from_tag $1 "genesis_with_funds"
}

function delegate_stake(){
  echo "delegate_stake for $1"
  local NODE_HOME="$1"
  local MONIKER=$(home_name "$NODE_HOME")
  each_create_gentx_for_delegation "$NODE_HOME" "$MONIKER" $VALIDATOR_INITIAL_BALANCE $CURRENCY
  each_write_delegation_gentx_to_repo "$NODE_HOME"
}

function each_create_gentx_for_delegation() {
  local HOME_DIR="${1:?"Home folder required"}"
  local MONIKER="${2:?"Moniker required"}"
  local AMOUNT=${3:?"Amount required"}
  local DENOM=${4:?"Currency required"}
  local KEYRING_BACKEND=${5:---keyring-backend test}
  local KEYRING_DIR=${6:---keyring-dir $HOME_DIR}
  local REPO="${HOME_DIR%/}/sync"
  local KEY_NAME=val-$MONIKER
  echo "$CHAIN_BINARY  --home $HOME_DIR gentx $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR $AMOUNT$DENOM --chain-id ${CHAIN_ID:?"CHAIN_ID required"} --ip $PUBLIC_HOST_NAME --moniker $MONIKER"
  $CHAIN_BINARY  --home $HOME_DIR gentx $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR $AMOUNT$DENOM --chain-id ${CHAIN_ID:?"CHAIN_ID required"} --ip $PUBLIC_HOST_NAME --moniker $MONIKER
}

function each_write_delegation_gentx_to_repo(){
  local HOME_DIR=${1:?"Home folder required"}
  local REPO="${HOME_DIR%/}/sync"
  local CONFIG="${HOME_DIR%/}/config"
  pull_git $REPO
  cp -r $CONFIG/gentx $REPO/config
  cd $REPO
  git add config
  git commit $GIT_QUIET -s -m "Add staking transaction for $MONIKER"
  push_git_retrying
  echo "Committed and pushed staking transaction for $MONIKER at $REPO/config/gentx"
  cd - > /dev/null
}

function fetch_genesis_file_with_txs(){
  echo "fetch_genesis_file_with_txs for $1"
  each_fetch_genesis_file_for_node_from_tag $1 "genesis_with_txs"
  adapt_app_toml $1
  adapt_client_toml $1
  adapt_config_toml $1
}

function adapt_app_toml(){
  echo "adapt_app_toml for $1"
  update_app_toml "${1%/}/sync" $1
}

function adapt_client_toml(){
  echo "adapt_client_toml for $1"
  update_client_toml "${1%/}/sync" $1
}

function adapt_config_toml(){
  echo "adapt_config_toml for $1"
  update_config_toml "${1%/}/sync" $1
}

function setup_node_main() {
  if [ -d $HOME_FOLDER ]; then
    exit 0;
  fi
  create_home_and_key $HOME_FOLDER
  each_write_address_to_repo $HOME_FOLDER
  fetch_genesis_file_with_funds $HOME_FOLDER
  delegate_stake $HOME_FOLDER
  fetch_genesis_file_with_txs $HOME_FOLDER
}

setup_node_main