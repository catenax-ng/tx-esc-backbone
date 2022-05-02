# Generate a set of validator home folders

[setup-nodes.sh](../../testnetscripts/setup-nodes.sh) used to prepare a folder structure and a docker-compose.yml to boot some nodes locally.

## Usage

```shell
GIT_DISCOVERY_ACROSS_FILESYSTEM=1 ./setup-nodes.sh <home-folders-parent> <orchestrator-home> \
 [<num-of-validators>] [<url-to-git-repo>]
```
## Parameters

### <home-folders-parent> 

### <orchestrator-home>
Mandatory parameter. folder of the component, which applies public address and genesis transactions into the genesis.json file.

### <num-of-validators>
Optional parameter. If omitted '4' is used.
Defines, how many home folder for validator nodes are 
Must be an integer x, where `x = 3 * n + 1` and `n` is a natural number.
The usage of docker-compose and IPv4 defines an upper bound of 253.
/!\ This constraint is not enforced yet.

### <url-to-git-repo>
Optional parameter. If omitted a temporary folder is used.
Pointing to a git repository used to gather public addresses and inital self delegation transactions. 

## Customizable Variables

If not otherwise stated, defined at `setup-nodes.sh`.

### NODE_NAME_PREFIX

Default `n`. 
The name of the validator node's home folder with its index as suffix.

### KEEP_LOCAL_REPO

Default `0`. 
If no `<url-to-git-repo>` is provided a temporary folder is used.
By setting this variable to `1` that temporary folder will not be deleted at the end of the script.

### DOCKER_COMPOSE_YAML_LOCATION

Default `<home-folders-parent>/../docker-compose.yml`. Defines the location, where the docker-compose.yml will be generated.

### NETWORK_ADDRESS_TEMPLATE

Default `172.16.0.{host}`. This will resolved with the validator nodes index to the ip address used in docker compose.

### NETWORK_ADDRESS

Default `172.16.0.0/24`. The docker-compose's network's address.

### NETWORK_GATEWAY_ADDRESS

Default `172.16.0.1`. The docker-compose's network's gateway/ host.

### CURRENCY

Default `ncax-demo`. Name of the currency used in the test net.

### CHAIN_ID

Default `catenax-testnet-1`. Id of the cosmos network.

### ADD_FAUCET_ACCOUNT

Default ` `. Set it to `i-know-this-is-insecure` to generate a faucet account. Mnemonic is given by `FAUCET_MNEMONIC`.


### FAUCET_MNEMONIC

Default `abuse submit area wide early west ripple oppose shed size describe foster need course lock use humble step film bridge timber unveil anxiety list`.
The mnemonic used to generate the key for the faucet's wallet.

### CHAIN_BINARY (at cosmos-helpers.sh)

Default `esc-backboned`. The name of the cosmos compatible node binary at the path. 

