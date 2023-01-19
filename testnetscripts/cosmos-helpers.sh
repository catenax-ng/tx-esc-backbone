#!/bin/bash

# Copyright (c) 2022 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

GIT_QUIET=" -q "
SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")
source $SCRIPT_LOCATION/helpers.sh
source $SCRIPT_LOCATION/update-app-toml.sh
source $SCRIPT_LOCATION/update-client-toml.sh
source $SCRIPT_LOCATION/update-config-toml.sh

ensure_command_exists $CHAIN_BINARY
ensure_command_exists pwd
ensure_command_exists grep
ensure_command_exists tail
ensure_command_exists sed
ensure_command_exists cat
ensure_command_exists touch
ensure_command_exists jq
# required to edit toml files
if ! command -v dasel &> /dev/null
then
  echo "dasel could not be found: can be installed with 'go install github.com/tomwright/dasel/cmd/dasel@v1.26.1'"
  exit
fi



function reset_validator_dir() {
  local HOME_DIR=${1:?"Home folder required"}
  $CHAIN_BINARY --home $HOME_DIR tendermint unsafe-reset-all
  rm $HOME_DIR/config/write-file-atomic-*
}
function add_funds_to_addr(){
  local HOME_DIR=${1:?"Home folder required"}
  local PUB_ADDR=${2:?"Public address required"}
  local AMOUNT=${3:?"Amount required"}
  local DENOM=${4:?"Currency required"}
  echo "$CHAIN_BINARY --home $HOME_DIR add-genesis-account $PUB_ADDR $AMOUNT$DENOM"
  $CHAIN_BINARY --home $HOME_DIR add-genesis-account $PUB_ADDR $AMOUNT$DENOM
}

function each_add_key(){
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=${2:?"Moniker required"}
  local _MNEMONIC="$3"
  local KEY_NAME=val-$MONIKER
  local KEYRING_BACKEND=${4:---keyring-backend test}
  local KEYRING_DIR=${5:---keyring-dir $HOME_DIR}
  if [ -z "$_MNEMONIC" ]; then
    $CHAIN_BINARY --home $HOME_DIR keys add $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR  |& grep -A 3 "It is the only way to recover your account if you ever forget your password." | tail -n +3 - > "$HOME_DIR"mnemonic-$KEY_NAME
  else
    echo "$_MNEMONIC" > "${HOME_DIR%/}/mnemonic-$KEY_NAME"
    echo "$_MNEMONIC" | $CHAIN_BINARY --home $HOME_DIR keys add $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR --recover
  fi
}


function get_pub_addr(){
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=${2:?"Moniker required"}
  local KEYRING_BACKEND=${3:---keyring-backend test}
  local KEYRING_DIR=${4:---keyring-dir $HOME_DIR}
  local KEY_NAME=val-$MONIKER
  echo $($CHAIN_BINARY --home $HOME_DIR keys show $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR -a)
}

function clone_sync_repo(){
  local REPO_TARGET=${1:?"target folder for git repo cloning required"}
  local _GIT_REPO=${2:?"git repo url for git repo cloning required"}
  local BRANCH="-B ${REPO_BRANCH:-main}"
  local retry=1
#  mkdir -p "$REPO_TARGET"
#  ls -la "$REPO_TARGET"
  git clone $GIT_QUIET "$_GIT_REPO" "$REPO_TARGET" > /dev/null
  local clone_result=$?
  while [ $clone_result -ne 0 -a $retry -lt ${GIT_WAIT_MAX_RETRY:-5} ]
  do
    ls -la $(dirname $REPO_TARGET)
    echo "Cloning repository $_GIT_REPO failed"
    echo "Retrying in ${GIT_WAIT:-1}s ... "
    sleep ${GIT_WAIT:-1}
    retry=$(( $retry + 1 ))
    git clone $GIT_QUIET "$_GIT_REPO" "$REPO_TARGET" > /dev/null
    local clone_result=$?
  done
  if [ $clone_result -ne 0 -a $retry -eq ${GIT_WAIT_MAX_RETRY:-5} ]
  then
    echo "Cloning repository $_GIT_REPO failed"
    echo "retries exceeded"
    exit 1
  fi
  cd $REPO_TARGET
  retry=1
  git checkout $GIT_QUIET ${BRANCH}
  local checkout_result=$?
  while [ $checkout_result -ne 0 -a $retry -lt ${GIT_WAIT_MAX_RETRY:-5} ]
  do
    echo "Checking out ${REPO_BRANCH:-main} from $_GIT_REPO after cloning failed failed"
    echo "Retrying in ${GIT_WAIT:-1}s ... "
    sleep ${GIT_WAIT:-1}
    retry=$(( $retry + 1 ))
    git clone $GIT_QUIET "$_GIT_REPO" "$REPO_TARGET" > /dev/null
    local checkout_result=$?
  done
  cd - > /dev/null
  if [ $checkout_result -ne 0 -a $retry -eq ${GIT_WAIT_MAX_RETRY:-5} ]
  then
    echo "Checking out ${REPO_BRANCH:-main} from $_GIT_REPO after cloning failed failed"
    echo "retries exceeded"
    exit 1
  fi
}


function sync_generate_genesis_file() {
  local TEMP_HOME=$(mktemp -d)
  $CHAIN_BINARY --home $TEMP_HOME init git-moniker --chain-id ${CHAIN_ID:?"CHAIN_ID required"} &> /dev/null
  local HOME_DIR=${1:?"Home folder required"}
  mkdir -p $HOME_DIR/config
  mv $TEMP_HOME/* $HOME_DIR/
  sed -i "s/\\\"stake\\\"/\"$CURRENCY\"/g" $HOME_DIR/config/genesis.json
  rm -r $TEMP_HOME
  tag_genesis_file $HOME_DIR "init_genesis" "initial genesis file"
}

function tag_genesis_file(){
  local REPO="${1:?'Git repo folder required'}"
  local TAG_NAME="${2:?'Tag name required'}"
  local COMMIT_MESSAGE="${3:?'Commit message required'}"
  echo "Tag $REPO with $TAG_NAME"
  cd $REPO
  if [ -f config/app.toml ]; then
    git add config/app.toml
  fi
  if [ -f config/client.toml ]; then
    git add config/client.toml
  fi
  if [ -f config/config.toml ]; then
    git add config/config.toml
  fi
  git add config/genesis.json
  git commit $GIT_QUIET -s -m "$COMMIT_MESSAGE"
  git tag -a $TAG_NAME -m "$COMMIT_MESSAGE"
  git push $GIT_QUIET --tags origin $REPO_BRANCH
  cd - > /dev/null
}

function pull_git() {
  local REPO=${1:?"Git repo folder required"}
  local TAG="$2"
  local BRANCH=${REPO_BRANCH:-main}
  cd $REPO
  git fetch $GIT_QUIET --tags
  if [ ! -z "$TAG" ]; then
    git checkout $GIT_QUIET tags/$TAG -b $TAG
  else
    git switch $GIT_QUIET ${BRANCH}
    git pull $GIT_QUIET origin  --ff-only ${BRANCH}
    local pull_result=$?
    if [  $pull_result -ne 0 ] ; then
      echo "Failed pulling fast forward $REPO"
      exit 1
    fi
  fi
  cd - > /dev/null
}


function wait_for_validator_commits() {
  local REPO=$1
  local GLOB="$2"
  local WAIT_FOR_NUM=$3
  pull_git $REPO
  echo "Wait for $WAIT_FOR_NUM validators to commit for $GLOB"
  shopt -s nullglob
  local VALIDATOR_INPUT=($GLOB)
  echo ${VALIDATOR_INPUT[@]}
  local retry=1
  local max_retry=$(( ${GIT_WAIT_MAX_RETRY:-5} * $WAIT_FOR_NUM ))
  while [ ${#VALIDATOR_INPUT[@]} -lt $WAIT_FOR_NUM -a $retry -lt $max_retry ]
  do
    echo "Not all validators committed yet (${#VALIDATOR_INPUT[@]}/$WAIT_FOR_NUM)"
    sleep ${GIT_WAIT:-1}
    retry=$(( $retry + 1 ))
    pull_git $REPO
    VALIDATOR_INPUT=($GLOB)
    echo ${VALIDATOR_INPUT[@]}
  done
  if [ ${#VALIDATOR_INPUT[@]} -lt $WAIT_FOR_NUM -a $retry -eq $max_retry ]
  then
      echo "Failed waiting for $WAIT_FOR_NUM validator commits at $GLOB with $retry retries."
      echo "This might be caused, if the number of validators is less than the value of \$VALIDATOR_COUNT (default: 4)"
      exit 2
  fi
  shopt -u nullglob
  echo "All validators committed"
}


