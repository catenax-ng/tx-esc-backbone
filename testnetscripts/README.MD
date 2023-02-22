<!--
 Copyright (c) 2022-2023 - for information on the respective copyright owner
 see the NOTICE file and/or the repository at
 https://github.com/catenax-ng/product-esc-backbone-code

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