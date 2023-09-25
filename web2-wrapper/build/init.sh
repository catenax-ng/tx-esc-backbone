#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0
KEY_NAME=$1
NATS_HOST=$2
cd config/
if [ -f mnemonic -a -f pubaddr ]; then
  /wrapper/esc-backboned query bank balances --home $(pwd) --node https://validator2-tdmt-rpc.dev.demo.catena-x.net:443/  $(cat pubaddr)
  exit $?
fi
cp ../default-config.json config.json
echo "Create mnemonic, this is to be kept secret, from this the private key can be generated."
/wrapper/esc-backboned keys mnemonic --home $(pwd) > mnemonic
echo "Import private key with mnemonic as 'wrapper'"
cat mnemonic | /wrapper/esc-backboned keys add $KEY_NAME --keyring-backend test --home $(pwd) --recover
echo "Store public address for the faucet request."
/wrapper/esc-backboned keys show $KEY_NAME --home $(pwd) --keyring-backend test -a > pubaddr
echo "Fetch the denom for currency from testnet"
DENOM=$(curl -s https://validator-webapp1-web-app.dev.demo.catena-x.net/chain/catenax-testnet-1-suggestion.json | jq ".feeCurrencies[0].coinMinimalDenom" -r)
echo "Request funds for $(cat pubaddr)"
curl -s -X POST --header "Content-Type: application/json" -d '{"address":"'$(cat pubaddr)'","denom":"'$DENOM'"}'  https://faucet-faucet.dev.demo.catena-x.net/
echo "Fetching current block height"
BLOCKHEIGHT=$(/wrapper/esc-backboned  status --node https://validator2-tdmt-rpc.dev.demo.catena-x.net:443/ --home $(pwd) | jq ".SyncInfo.latest_block_height" -r)
echo "Replace nats url ($NATS_HOST), fee currency ($DENOM), block height ($BLOCKHEIGHT) and keyname ($KEY_NAME) in config.json"
echo $(jq "setpath(path(.from);\"$KEY_NAME\")|setpath(path(.broker.url);\"$NATS_HOST\")|setpath(path(.fees);\"10000$DENOM\")|setpath(path(.start_block);$BLOCKHEIGHT)" config.json) > config.json
echo "Wait 10s and check the balance"
sleep 10 && /wrapper/esc-backboned query bank balances --home $(pwd) --node https://validator2-tdmt-rpc.dev.demo.catena-x.net:443/  $(cat pubaddr)