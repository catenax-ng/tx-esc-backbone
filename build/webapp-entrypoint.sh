#!/bin/bash
SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")


${SCRIPT_LOCATION%/}/update-chain-suggestion.sh ${CHAIN_SUGGESTION:-/usr/share/nginx/html/chain/catenax-testnet-1-suggestion.json}
# entrypoint of https://github.com/nginxinc/docker-nginx-unprivileged/
/docker-entrypoint.sh "$@"