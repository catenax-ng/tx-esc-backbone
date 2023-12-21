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


# This script reads the environment variables and writes the one with the prefix 'WEBAPP_' to the given JSON file.
# This enables the static web app to access these environment variables by doing a HTTP GET.
# The image works in Docker and Kubernetes by just passing environment variables to the deployed container.
# This omits the necessity to mount the config file, which is quite tool dependent.
JSON_FILE=${1:?"File name to store env variables required."}
PREFIX='WEBAPP_'

# start json object
echo "{" > $JSON_FILE
while IFS= read -r line; do
    echo $line | \
    # filter on variables with the prefix
    grep -e "^$PREFIX"  | \
    # skip non matching variables
    grep -E '^.+$' | \
    # transform variable into json attribute and write it to the file
    sed -r "s/^([^=\s]*)=(.*)/\"\1\":\"\2\",/g" >> $JSON_FILE
done <<< $(env)
# end json object
echo "}" >> $JSON_FILE

# remove comma after last variable
sed -i -r ":begin;$!N;s/,\n}/\n}/;tbegin;P;D" $JSON_FILE