#!/bin/bash

# Copyright (c) 2022 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

JSON_FILE=${1:?"File name to keplr chain suggestion required."}
JQ_PROGRAM=""

if [ ! -z $CHAIN_REST ]; then
  if [ ! -z $JQ_PROGRAM ]; then JQ_PROGRAM="$JQ_PROGRAM|"; fi
  JQ_PROGRAM="${JQ_PROGRAM}setpath(path(.rest);\"$CHAIN_REST\")"
fi

if [ ! -z $CHAIN_RPC ]; then
  if [ ! -z $JQ_PROGRAM ]; then JQ_PROGRAM="$JQ_PROGRAM|"; fi
  JQ_PROGRAM="${JQ_PROGRAM}setpath(path(.rpc);\"$CHAIN_RPC\")"
fi

if [ ! -z $CHAIN_ID ]; then
  if [ ! -z $JQ_PROGRAM ]; then JQ_PROGRAM="$JQ_PROGRAM|"; fi
  JQ_PROGRAM="${JQ_PROGRAM}setpath(path(.chainId);\"$CHAIN_ID\")"
fi

if [ ! -z $CHAIN_NAME ]; then
  if [ ! -z $JQ_PROGRAM ]; then JQ_PROGRAM="$JQ_PROGRAM|"; fi
  JQ_PROGRAM="${JQ_PROGRAM}setpath(path(.chainName);\"$CHAIN_NAME\")"
fi
if [ !  -z $CHAIN_DENOM_NEW  ] && [ ! -z $CHAIN_DENOM_FILE  ]; then
  echo "Replacing denom $CHAIN_DENOM_FILE with $CHAIN_DENOM_NEW"
  if [ ! -z $JQ_PROGRAM ]; then JQ_PROGRAM="$JQ_PROGRAM|"; fi
  # operate at the currencies and start an array
  JQ_PROGRAM="${JQ_PROGRAM}setpath(path(.currencies);["
  # select the currency with the denom $CHAIN_DENOM_FILE
  JQ_PROGRAM="${JQ_PROGRAM}.currencies[]|(select(.coinDenom==\"$CHAIN_DENOM_FILE\")"
  # set the denom at this currency to  $CHAIN_DENOM_NEW
  JQ_PROGRAM="${JQ_PROGRAM}|setpath(path(.coinDenom);\"$CHAIN_DENOM_NEW\"))"
  # add all unmodified currencies.
  JQ_PROGRAM="${JQ_PROGRAM},select(.coinDenom!=\"$CHAIN_DENOM_FILE\")"
  # finish the array
  JQ_PROGRAM="${JQ_PROGRAM}])"
  echo "$JQ_PROGRAM"
  # operate at the fee currencies and start an array
  JQ_PROGRAM="${JQ_PROGRAM}|setpath(path(.feeCurrencies);["
  # select the currency with the denom $CHAIN_DENOM_FILE
  JQ_PROGRAM="${JQ_PROGRAM}.feeCurrencies[]|(select(.coinDenom==\"$CHAIN_DENOM_FILE\")"
  # set the denom at this currency to  $CHAIN_DENOM_NEW
  JQ_PROGRAM="${JQ_PROGRAM}|setpath(path(.coinDenom);\"$CHAIN_DENOM_NEW\"))"
  # add all unmodified currencies.
  JQ_PROGRAM="${JQ_PROGRAM},select(.coinDenom!=\"$CHAIN_DENOM_FILE\")"
  # finish the array
  JQ_PROGRAM="${JQ_PROGRAM}])"
  echo "$JQ_PROGRAM"
  STAKE_CURRENCY=$(jq ".stakeCurrency.coinDenom" $JSON_FILE)
  echo $STAKE_CURRENCY
  echo $CHAIN_DENOM_FILE
  echo [ $STAKE_CURRENCY = $CHAIN_DENOM_FILE]
  if [ $STAKE_CURRENCY = "\"$CHAIN_DENOM_FILE\"" ]; then
    JQ_PROGRAM="${JQ_PROGRAM}|setpath(path(.stakeCurrency.coinDenom);\"$CHAIN_DENOM_NEW\")"
    echo "$JQ_PROGRAM"
  fi
fi

if [ ! -z $JQ_PROGRAM ]; then
  echo $(jq "${JQ_PROGRAM}" $JSON_FILE) > $JSON_FILE
fi