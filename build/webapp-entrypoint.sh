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


SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")

${SCRIPT_LOCATION%/}/write-env-to-json.sh ${ENV_JSON:-/usr/share/nginx/html/chain/env.json}
${SCRIPT_LOCATION%/}/update-chain-suggestion.sh ${CHAIN_SUGGESTION:-/usr/share/nginx/html/chain/catenax-testnet-1-suggestion.json}
# entrypoint of https://github.com/nginxinc/docker-nginx-unprivileged/
/docker-entrypoint.sh "$@"