#!/bin/bash
# Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
#
# See the NOTICE file(s) distributed with this work for additional
# information regarding copyright ownership.
#
# This program and the accompanying materials are made available under the
# terms of the Apache License, Version 2.0 which is available at
# https://www.apache.org/licenses/LICENSE-2.0.
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.
#
# SPDX-License-Identifier: Apache-2.0


if [ -z "$CHAIN_BINARY" ]; then
  export CHAIN_BINARY=esc-backboned
fi

if [ -z "$REPO_BRANCH" ]; then
  export REPO_BRANCH=main
fi

if [ -z "$CHAIN_ID" ]; then
  export CHAIN_ID="catenax-testnet-1"
fi

if [ -z "$CURRENCY" ]; then
  export CURRENCY="ncaxdemo"
fi

if [ -z "$GIT_WAIT" ]; then
  export GIT_WAIT=3
fi

if [ -z "$VALIDATOR_COUNT" ]; then
  export VALIDATOR_COUNT=4
fi

if [ -z "$GIT_WAIT_MAX_RETRY" ]; then
  export GIT_WAIT_MAX_RETRY=20
fi

if [ -z "$GIT_PUSH_MAX_RETRY" ]; then
  export GIT_PUSH_MAX_RETRY=$(( $VALIDATOR_COUNT + 2 ))
fi

if [ -z "$FAUCET_INITIAL_BALANCE" ]; then
  export FAUCET_INITIAL_BALANCE="1000000000000000000000"
fi

if [ -z "$VALIDATOR_INITIAL_BALANCE" ]; then
  export VALIDATOR_INITIAL_BALANCE="1000000000000000000000"
fi

if [ -z "$FAUCET_MNEMONIC" ]; then
  export FAUCET_MNEMONIC="abuse submit area wide early west ripple oppose shed size describe foster need course lock use humble step film bridge timber unveil anxiety list"
fi
# export ADD_FAUCET_ACCOUNT="i-know-this-is-insecure"

if [ ! -z "$GIT_REPO" ]; then
  # resolve variables in single quotes variable values
  export GIT_REPO=$(eval "echo $GIT_REPO")
fi