#!/bin/bash


SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}" )
source $SCRIPT_LOCATION/cosmos-helpers.sh

NODES="${1:-work}"
if [ ! -d "$NODES" ]; then
  echo "$NODES is not a folder. Aborting" >&2
  exit 1
fi

main() {
  NODE_HOMES=($NODES/*/)
  apply_on_each reset_validator_dir ${NODE_HOMES[@]}
}

main
