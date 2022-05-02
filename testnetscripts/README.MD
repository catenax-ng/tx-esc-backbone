# Generate a set of validator home folders
[setup-nodes.sh](./setup-nodes.sh) used to prepare a folder structure and a docker-compose.yml to boot some nodes locally.

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