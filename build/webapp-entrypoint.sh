#!/bin/bash

# Copyright (c) 2022 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")

${SCRIPT_LOCATION%/}/write-env-to-json.sh ${ENV_JSON:-/usr/share/nginx/html/chain/env.json}
${SCRIPT_LOCATION%/}/update-chain-suggestion.sh ${CHAIN_SUGGESTION:-/usr/share/nginx/html/chain/catenax-testnet-1-suggestion.json}
# entrypoint of https://github.com/nginxinc/docker-nginx-unprivileged/
/docker-entrypoint.sh "$@"