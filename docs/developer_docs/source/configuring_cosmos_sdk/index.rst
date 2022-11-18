.. Copyright (c) 2022 - for information on the respective copyright owner
.. see the NOTICE file and/or the repository at
.. https://github.com/catenax-ng/product-esc-backbone-code
..
.. SPDX-License-Identifier: Apache-2.0

.. _configuring_cosmos_sdk:

Configuring the cosmos sdk
==========================

This section provides an overview what aspects of the node can be configured,
how and where to do it. It is intended as a starting point for new developers
joining the team and as a quick reference for all developers.

Configuration files
-------------------

The configuration files are stored in the `config` directory in the node's home
folder. This directory consists of 3 configuration files:

1. `app.toml`: Configures the aspects related to application such as REST API,
   gRPC API, Rosetta, Telemetry etc.,.

2. `config.toml`: Configures the aspects related to tendermint node, such as
   genesis file, validator and node key files, state db, tendermint RPC, P2P,
   mempool, state sync, consensus and some other aspects.

3. `genesis.json`: Configures the initial parameters of the blockchain. It
   includes allowed pub key types and parameters for each included modules.


Different endpoints exposed by a running node
---------------------------------------------

There are 3 different set of APIs, that are relevant for the user of the node.
These are

1. gRPC API 

   - This is exposed by blockchain application and default port is 9090.
   - This is configured in `app.toml`.
   - The services for querying a module and for broadcasting signed
     transactions can be accessed via this interface.
   - See `this section <https://docs.cosmos.network/main/core/grpc_rest.html>`_
     of cosmos sdk docs for further info.

2. REST API

   - This is exposed by blockchain application and default port is 1317.
   - This is configured in `app.toml`.
   - The gRPC endpoints can be accessed as REST via this interface.
   - Internally, it is implemented using a gRPC gateway.
   - See `this section <https://docs.cosmos.network/main/core/grpc_rest.html>`_
     of cosmos sdk docs for further info.

3. Tendermint RPC API (default port: 26657)

   - This is exposed directly by the tendermint node and default port is 26657.
   - This is configured in `config.toml`.
   - All the endpoints of tendermint RPC are published as OpenAPI specification
     `here <https://docs.tendermint.com/master/rpc/>`_.
   - Other APIs (gRPC and REST), use the tendermint API under the hood for
     broadcasting transactions.
   - See `this section <https://docs.cosmos.network/main/core/grpc_rest.html>`_
     of cosmos sdk docs for further info.


Additionally, there are two other endpoints. These are

1. Tendermint P2P

   - This is exposed directly by the tendermint and default port is 26656.
   - This is configured in `config.toml`.
   - It is used for peer to peer interaction with other nodes in the blockchain
     network.
   - See `this section
     <https://docs.tendermint.com/v0.33/tendermint-core/secure-p2p.html>`_ of
     tendermint docs for further info.
   

2. Prometheus

   - This is exposed directly by the tendermint node and default port is 26660.
   - This is configured in `config.toml`.
   - It reports and serves prometheus metrics.
   - See `this section
     <https://docs.tendermint.com/master/nodes/metrics.html#metrics>`_ of
     tendermint docs for further info.


Keys
----

In the cosmos SDk, two types of keys are used:

1. Application keys: Used for signing transactions created by the user of this
   node.

2. Tendermint keys: Used for signing votes, as a part of the validation
   process.

Modules
-------

The architecture of the node consists of the two components:

1. Blockchain application: It interprets, validates and maintains the state;
   implements proof of stake consensus over tendermint consensus.

2. Tendermint node: It syncs the state across all the nodes in a bynzatine
   fault tolerant manner, using the tendermint consensus mechanism.


The blockchain application is composed of many modules, each of which provides
certain functionalities. These modules can be existing ones created by the core
team, community or the ones developed specifically for our application. These
modules are configured by modifying their parameters in the genesis file.

Below is a set of modules included in the esc-backbone blockchain node, a brief
description of their functionality and their parameters. See `this page
<https://docs.cosmos.network/main/modules/>`_ for more information on any of
these modules.

1. auth: accounts, signatures, transactions and gas.

   Note: one aspect of gas cost "constant_fee" is configured in the crisis
   module.

2. genutil: genesis utility for use within the application. 

   It has no configurable parameters.

3. bank module: supply and minting/burning of tokens.

4. capability: implements object capability, as described in `ADR 3
   <https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-003-dynamic-capability-store.md>`_.

5. staking: staking and proof-of-stake (PoS) consensus.

6. mint: implements inflation

7. distribution: passive mechanism to manage rewards and community tax.

8. governance: on-chain governance proposals for modifying parameters, software
   upgrades and general purpose concerns.

9. params: global aware parameter store used by all other modules.

10. crisis: ensures chain is halted when invariants are broken.

    It has a configurable parameter call "constant fee" which is the constant
    fee charged for checking the invariants.

11. feegrant: allow accounts to grant fee allowances and use fees from their
    accounts.

12. IBC: implements inter-blockchain communication protocol.

13. upgrade: for smooth upgrade of live chain to breaking software version.

14. evidence: allow submission of arbitrary evidence of misbehavior.

15. transfer: transfer of tokens across networks via IBC.

16. vesting: vesting of tokens.
