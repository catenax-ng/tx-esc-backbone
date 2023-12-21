<!--
 Copyright (c) 2022-2023 Contributors to the Eclipse Foundation

 See the NOTICE file(s) distributed with this work for additional
 information regarding copyright ownership.

 This program and the accompanying materials are made available under the
 terms of the Apache License, Version 2.0 which is available at
 https://www.apache.org/licenses/LICENSE-2.0.

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 License for the specific language governing permissions and limitations
 under the License.

 SPDX-License-Identifier: Apache-2.0
-->

# Generate a set of validator home folders
[setup-nodes.sh](./setup-nodes.sh) used to prepare a folder structure to boot some nodes.

## Usage
```shell
GIT_DISCOVERY_ACROSS_FILESYSTEM=1 ./setup-nodes.sh <home-folders-parent> <orchestrator-home> \
 [<num-of-validators>] [<url-to-git-repo>]
```


Example, which was used to generate the `testnet` folder in this repository
```shell
rm -r ./testnet && GIT_DISCOVERY_ACROSS_FILESYSTEM=1 ADD_FAUCET_ACCOUNT=i-know-this-is-insecure ./testnetscripts/setup-nodes.sh ./testnet/work ./testnet/orchestrator 4
```

## Details
See [generate_testnet_homes.md](../docs/testnet_homes/generate_testnet_homes.md).