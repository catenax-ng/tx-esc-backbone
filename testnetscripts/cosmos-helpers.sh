#!/bin/bash
GIT_QUIET=" -q "
SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")
source $SCRIPT_LOCATION/helpers.sh

CHAIN_BINARY=esc-backboned

ensure_command_exists $CHAIN_BINARY
ensure_command_exists pwd
ensure_command_exists grep
ensure_command_exists tail

function each_init_home(){
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=${2:?"Moniker required"}
  $CHAIN_BINARY --home $HOME_DIR init $MONIKER &> /dev/null
}

function add_funds_to_addr(){
  local HOME_DIR=${1:?"Home folder required"}
  local PUB_ADDR=${2:?"Public address required"}
  local AMOUNT=${3:?"Amount required"}
  local DENOM=${4:?"Currency required"}
  echo "$CHAIN_BINARY --home $HOME_DIR add-genesis-account $PUB_ADDR $AMOUNT$DENOM"
  $CHAIN_BINARY --home $HOME_DIR add-genesis-account $PUB_ADDR $AMOUNT$DENOM
}

function each_create_gentx_for_delegation() {
  local HOME_DIR="${1:?"Home folder required"}"
  local MONIKER="${2:?"Moniker required"}"
  local AMOUNT=${3:?"Amount required"}
  local DENOM=${4:?"Currency required"}
  local KEYRING_BACKEND=${5:---keyring-backend test}
  local KEYRING_DIR=${6:---keyring-dir $HOME_DIR}
  local REPO="${HOME_DIR%/}/sync"
  local IP_ADDRESS=$(cat ${HOME_DIR%/}/ip_address)
  local KEY_NAME=val-$MONIKER
  echo "$CHAIN_BINARY  --home $HOME_DIR gentx $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR $AMOUNT$DENOM --chain-id ${CHAIN_ID:?"CHAIN_ID required"} --ip $IP_ADDRESS --moniker $MONIKER"
  $CHAIN_BINARY  --home $HOME_DIR gentx $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR $AMOUNT$DENOM --chain-id ${CHAIN_ID:?"CHAIN_ID required"} --ip $IP_ADDRESS --moniker $MONIKER
}

function each_add_key(){
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=${2:?"Moniker required"}
  local MNEMONIC=$3
  local KEY_NAME=val-$MONIKER
  local KEYRING_BACKEND=${4:---keyring-backend test}
  local KEYRING_DIR=${5:---keyring-dir $HOME_DIR}
  if [ -z "$MNEMONIC" ]; then
    $CHAIN_BINARY --home $HOME_DIR keys add $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR  |& grep -A 3 "It is the only way to recover your account if you ever forget your password." | tail -n +3 - > "$HOME_DIR"mnemonic-$KEY_NAME
  else
    echo "recover"
    cat $MNEMONIC | $CHAIN_BINARY --home $HOME_DIR keys add $KEY_NAME $KEYRING_BACKEND $KEYRING_DIR --recover
  fi
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
  git push $GIT_QUIET
  cd - > /dev/null
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
  local REPO_URL=${2:?"git repo url for git repo cloning required"}
  local REPO_TARGET_PARENT=$(dirname $SCRIPT_LOCATION/REPO_TARGET)
  local BRANCH=-B main
  if [ ! -z "$REPO_BRANCH" ]; then
    BRANCH="-B $REPO_BRANCH"
  fi
  git clone $GIT_QUIET "$REPO_URL" "$REPO_TARGET" > /dev/null
  cd $REPO_TARGET
  git checkout $GIT_QUIET ${BRANCH}
  cd - > /dev/null
}

function each_fetch_genesis_file_for_node_from_tag(){
  local HOME_DIR=${1:?"Home folder required"}
  local REPO="${HOME_DIR%/}/sync"
  pull_git $REPO ${2:?"tag name required"}
  mkdir -p "${HOME_DIR%/}/config/"
  cp "$REPO/config/genesis.json" "${HOME_DIR%/}/config/"
}

function each_write_address_to_repo(){
  local HOME_DIR=${1:?"Home folder required"}
  local MONIKER=$(home_name "$HOME_DIR")
  local REPO="${HOME_DIR%/}/sync"
  pull_git $REPO
  mkdir -p $REPO/pub_addr
  local PUB_ADDR=$(get_pub_addr "$HOME_DIR" "$MONIKER")
  touch $REPO/pub_addr/$PUB_ADDR
  echo "$MONIKER" > $REPO/pub_addr/$PUB_ADDR
  cd $REPO
  git add pub_addr/$PUB_ADDR
  git commit $GIT_QUIET -s -m "Add public address for $MONIKER"
  git push $GIT_QUIET
  cd - > /dev/null
}

function sync_generate_genesis_file() {
  local HOME_DIR=${1:?"Home folder required"}
  local TEMP_HOME=$(mktemp -d)
  $CHAIN_BINARY --home $TEMP_HOME init git-moniker --chain-id ${CHAIN_ID:?"CHAIN_ID required"} &> /dev/null
  local HOME_DIR=${1:?"Git repo folder required"}
  mkdir -p $HOME_DIR/config
  mv $TEMP_HOME/config/genesis.json $HOME_DIR/config/
  sed -i "s/\\\"stake\\\"/\"$CURRENCY\"/g" $HOME_DIR/config/genesis.json
  rm -r $TEMP_HOME
  tag_genesis_file $HOME_DIR "init_genesis" "initial genesis file"
}

function sync_apply_gentxs(){
    local HOME_DIR=${1:?"Home folder required"}
    pull_git "$HOME_DIR"
    echo "$CHAIN_BINARY --home $HOME_DIR collect-gentxs"
    $CHAIN_BINARY --home "$HOME_DIR" collect-gentxs
    tag_genesis_file $HOME_DIR "genesis_with_txs" "genesis file with transactions"
}

function tag_genesis_file(){
  local REPO="${1:?'Git repo folder required'}"
  local TAG_NAME="${2:?'Tag name required'}"
  local COMMIT_MESSAGE="${3:?'Commit message required'}"
  cd $REPO
  git add config/genesis.json
  git commit $GIT_QUIET -s -m "$COMMIT_MESSAGE"
  git tag -a $TAG_NAME -m "$COMMIT_MESSAGE"
  git push $GIT_QUIET --tags origin $REPO_BRANCH
  cd - > /dev/null
}

function pull_git() {
  local REPO=${1:?"Git repo folder required"}
  local TAG="$2"
  cd $REPO
  git fetch $GIT_QUIET --tags
  if [ ! -z "$TAG" ]; then
    git checkout $GIT_QUIET tags/$TAG -b $TAG
  else
    git switch $GIT_QUIET ${REPO_BRANCH:-main}
    git pull $GIT_QUIET origin ${REPO_BRANCH:-main}
  fi
  cd - > /dev/null
}


