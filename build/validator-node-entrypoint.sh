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


id
MOUNT_FOLDER=${1:?"Mount folder required"}
MOUNT_FOLDER=${MOUNT_FOLDER%/}
if [ ! -w $MOUNT_FOLDER ]; then
  echo "Cannot write to $MOUNT_FOLDER"
  exit 1 
fi
HOME_NAME=${2:?"Provide the name of the home folder"}
HOME_FOLDER="$MOUNT_FOLDER/$HOME_NAME"
TEMPLATE_FOLDER=${3%/}
if [ ! -z $TEMPLATE_FOLDER ]; then
  if [ ! -d $HOME_FOLDER ]; then
    echo "Copy $TEMPLATE_FOLDER to $HOME_FOLDER"
    mkdir -p "$HOME_FOLDER"
    cp -r $TEMPLATE_FOLDER/* "$HOME_FOLDER"
  fi
fi

if [ ! -d "$HOME_FOLDER" ]; then
  echo "$HOME_FOLDER does not exist."
  exit 1
fi

echo "Starting /validator/esc-backboned with home: $HOME_FOLDER"
/validator/esc-backboned start --home "$HOME_FOLDER"
