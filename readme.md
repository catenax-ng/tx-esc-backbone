<!--
Copyright (c) 2022 - for information on the respective copyright owner
see the NOTICE file and/or the repository at
https://github.com/catenax-ng/product-esc-backbone-code

SPDX-License-Identifier: Apache-2.0
-->

# ESC Backbone
**ESC Backbone** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

The web frontend is a React based application that has a faucet client and
serves the chain suggestion for Keplr wallet.
Run the following commands to install dependencies and start the app:

```
cd web
npm install
npm run build && npx http-server dist
```

### Install
To install the latest version of the blockchain node's binary, execute the following command on your machine:

```
git clone https://github.com/catenax-ng/product-esc-backbone-code.git
ignite chain build
```

`esc-backboned` binary is created in the $GOPATH/bin directory.

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
