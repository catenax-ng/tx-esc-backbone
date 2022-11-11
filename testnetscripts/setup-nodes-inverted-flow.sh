#!/bin/bash
NODE_NAME_PREFIX=n
KEEP_LOCAL_REPO=0
NETWORK_ADDRESS_TEMPLATE="testnet-esc-backbone-{host}-service"

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/cosmos-helpers.sh

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

export GIT_REPO="${4}"
if [ -z "$GIT_REPO" ]; then
 export GIT_REPO=$(mktemp -d)
 GIT_TEMP_REPO="$GIT_REPO"
 echo "Using the temporary folder $GIT_REPO as remote repository."
 echo "To keep it, set KEEP_LOCAL_REPO=1."
 echo "If this script fails run 'rm -rf $GIT_REPO' to clean the remainder."
 create_a_local_empty_repo $GIT_REPO $REPO_BRANCH
fi

echo "generating homes for $NODE_COUNT nodes at $NODES with repo at $GIT_REPO."

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
  $SCRIPT_LOCATION/orchestrator-network-setup.sh $SYNC_HOME &
  ORCH_PID=$!
  # collect node homes
  NODE_HOMES=($NODES/*/)

  # start setup_node.sh for each home and wait
  for entry in "${NODE_HOMES[@]}"
  do
    local MNEMONIC=$($CHAIN_BINARY keys mnemonic)
    local PUBLIC_HOST_NAME=$(cat ${entry%/}/ip_address)
    $SCRIPT_LOCATION/setup-node.sh $entry $PUBLIC_HOST_NAME $MNEMONIC &
    local pid=$!
    declare node_pid$(basename $entry)=$pid
  done
  for entry in "${NODE_HOMES[@]}"
  do
    local node_pid_var_name=node_pid$(basename $entry)
    local node_pid=${!node_pid_var_name}
    wait ${node_pid}
  done
  wait $ORCH_PID

  cleanup_temp_repo
}

main
