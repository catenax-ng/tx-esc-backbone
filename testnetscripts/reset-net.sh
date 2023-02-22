#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0


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
