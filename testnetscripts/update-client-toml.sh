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


function update_client_toml(){
  local SRC=${1:?"Source folder required"}
  local TRG=${2:?"Target folder required"}
  # dasel strips comments for now. https://github.com/TomWright/dasel/issues/178
  # creating a backup
  local CONFIG_FILE_NAME="config/client.toml"
  local TRG_FILE="${TRG%/}/$CONFIG_FILE_NAME"
  local SRC_FILE="${SRC%/}/$CONFIG_FILE_NAME"
  cp "${TRG_FILE}" "${TRG_FILE}.bak"
  echo "cp "${TRG_FILE}" "${TRG_FILE}.bak""
}