#!/bin/bash

REPO_BRANCH=main
NODE_NAME_PREFIX=n
KEEP_LOCAL_REPO=0
NETWORK_ADDRESS_TEMPLATE="testnet-esc-backbone-{host}-service"
NETWORK_ADDRESS="172.16.0.0/24"
NETWORK_GATEWAY_ADDRESS="172.16.0.1"
DOCKER_COMPOSE_YAML_LOCATION=
CURRENCY="ncax-demo"
CHAIN_ID="catenax-testnet-1"
# ADD_FAUCET_ACCOUNT="i-know-this-is-insecure"
FAUCET_MNEMONIC=${FAUCET_MNEMONIC:-"abuse submit area wide early west ripple oppose shed size describe foster need course lock use humble step film bridge timber unveil anxiety list"}
FAUCET_INITIAL_BALANCE="1000000000000000000000"

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/cosmos-helpers.sh
source $SCRIPT_LOCATION/docker-compose-templating/docker-compose-template.sh

ensure_command_exists mktemp
ensure_command_exists cat
ensure_command_exists basename
ensure_command_exists git

NODES="${1:-work}"
if [ -e "$NODES" ]; then
  echo "$NODES already exists. Aborting" >&2
  exit 1
fi
mkdir -p "$NODES"

SYNC_HOME="${2:-genesis-sync}"
NODE_COUNT=${3:-4}
if ! is_unsigned_int $NODE_COUNT ; then
  NODE_COUNT=4
fi

GIT_REPO="${4}"
if [ -z "$GIT_REPO" ]; then
 GIT_REPO=$(mktemp -d)
 GIT_TEMP_REPO="$GIT_REPO"
 echo "Using the temporary folder $GIT_REPO as remote repository."
 echo "To keep it, set KEEP_LOCAL_REPO=1."
 echo "If this script fails run 'rm -rf $GIT_REPO' to clean the remainder."
 create_a_local_empty_repo $GIT_REPO $REPO_BRANCH
fi

echo "generating homes for $NODE_COUNT nodes at $NODES with repo at $GIT_REPO."

if [ -z "$DOCKER_COMPOSE_YAML_LOCATION"]; then
  DOCKER_COMPOSE_YAML_LOCATION="$NODES/../docker-compose.yml"
  echo "docker-compose.yml will be generated at $DOCKER_COMPOSE_YAML_LOCATION"
fi

function generate_home_folders_for_nodes() {
  echo "generate_home_folders_for_nodes with $NODE_COUNT"
  for ((i=1; i<=$NODE_COUNT; i++)); do
    echo "Generating $NODE_NAME_PREFIX$i"
    local NODE_HOME=$NODES/$NODE_NAME_PREFIX$i
    mkdir -p "$NODE_HOME"
    let ip=$i
    echo $NETWORK_ADDRESS_TEMPLATE | sed "s/{host}/$ip/g" > "$NODE_HOME/ip_address"
    echo $i > "$NODE_HOME/node_index"
  done
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
  rm "$NODE_HOME/config/genesis.json"
  each_fetch_genesis_file_for_node_from_tag "$NODE_HOME" "init_genesis"
}

function publish_initial_genesis_file() {
  echo "publish_initial_genesis_file at $SYNC_HOME"
  clone_sync_repo "$SYNC_HOME" "$GIT_REPO"
  sync_generate_genesis_file "$SYNC_HOME"
}

function add_funds_for_initial_account() {
  echo "add_funds_for_initial_account for $@"
  local PUB_ADDR_FILE=${1:?"Requires a file with the public address as name and the moniker as content."}
  local PUB_ADDR=$(basename $PUB_ADDR_FILE)
  local MONIKER=$(cat $PUB_ADDR_FILE)
  add_funds_to_addr "$SYNC_HOME" $PUB_ADDR "100000000000" $CURRENCY
}

function add_initial_accounts() {
  echo "add_faucet_account at $SYNC_HOME"
  pull_git "$SYNC_HOME"
  local PUB_ADDRS=($SYNC_HOME/pub_addr/*)
  apply_on_each add_funds_for_initial_account ${PUB_ADDRS[@]}
  add_faucet_account
  tag_genesis_file "$SYNC_HOME" "genesis_with_funds" "add funds to genesis file"
}

function add_faucet_account(){
  echo "add_faucet_account"
  if  [ "i-know-this-is-insecure" == "$ADD_FAUCET_ACCOUNT" ]; then
    echo "!!!Adding a faucet for mnemonic: $FAUCET_MNEMONIC"
    local FAUCET_DIR="${SYNC_HOME%/}/faucet"
    mkdir $FAUCET_DIR
    local MONIKER=$(home_name "$FAUCET_DIR")
    FAUCET_MNEMONIC_FILE=$FAUCET_DIR/mnemonic
    echo "$FAUCET_MNEMONIC" > $FAUCET_MNEMONIC_FILE
    each_add_key $FAUCET_DIR $MONIKER $FAUCET_MNEMONIC_FILE
    cd "${SYNC_HOME}"
    git add faucet
    git commit -m "Add faucet account"
    cd - > /dev/null
    local PUB_ADDR=$(get_pub_addr $FAUCET_DIR $MONIKER)
    add_funds_to_addr "$SYNC_HOME" $PUB_ADDR $FAUCET_INITIAL_BALANCE $CURRENCY
  fi
}

function delegate_stake(){
  echo "delegate_stake for $1"
  local NODE_HOME="$1"
  local MONIKER=$(home_name "$NODE_HOME")
  each_create_gentx_for_delegation "$NODE_HOME" "$MONIKER" "1000000000" $CURRENCY
  each_write_delegation_gentx_to_repo "$NODE_HOME"
}

function fetch_genesis_file_with_funds(){
  echo "fetch_genesis_file_with_funds for $1"
  each_fetch_genesis_file_for_node_from_tag $1 "genesis_with_funds"
}

function fetch_genesis_file_with_txs(){
  echo "fetch_genesis_file_with_txs for $1"
  each_fetch_genesis_file_for_node_from_tag $1 "genesis_with_txs"
  adapt_app_toml $1
  adapt_client_toml $1
  adapt_config_toml $1
}

function write_docker_compose_file() {
  echo "###TODO write_docker_compose_file"
}

function cleanup_temp_repo() {
  if [ ! -z $GIT_TEMP_REPO ]; then
    if [ $KEEP_LOCAL_REPO -ne 1 ]; then
      echo "rm -rf $GIT_TEMP_REPO"
      rm -rf $GIT_TEMP_REPO
    else
      echo "$GIT_TEMP_REPO will not be deleted"
    fi
  fi
}

main() {
  # generate node homes
  generate_home_folders_for_nodes
  # generate initial genesis file
  publish_initial_genesis_file
  # collect node homes
  NODE_HOMES=($NODES/*/)
  # each create a key
  apply_on_each create_home_and_key ${NODE_HOMES[@]}
  # each add their key
  apply_on_each each_write_address_to_repo ${NODE_HOMES[@]}
  # sync add funds and add tag
  add_initial_accounts
  # each load genesis file
  apply_on_each fetch_genesis_file_with_funds ${NODE_HOMES[@]}
  # each add delegation
  apply_on_each delegate_stake ${NODE_HOMES[@]}
  # sync apply gentxs and add tag
  sync_apply_gentxs $SYNC_HOME
  # each load genesis file and adapt *.toml files
  apply_on_each fetch_genesis_file_with_txs ${NODE_HOMES[@]}
  # write docker-compose.yaml
  generate_docker_compose_yml "$DOCKER_COMPOSE_YAML_LOCATION" $NODES $NETWORK_ADDRESS $NETWORK_GATEWAY_ADDRESS

  cleanup_temp_repo
}

main
