<!--
 Copyright (c) 2022-2023 - for information on the respective copyright owner
 see the NOTICE file and/or the repository at
 https://github.com/catenax-ng/product-esc-backbone-code

 SPDX-License-Identifier: Apache-2.0
-->

# ESC Backbone
**ESC Backbone** is a blockchain built using Cosmos SDK and Comet BFT and was created with [Ignite](https://ignite.com/cli).

## Maturity
This repository is not released yet.
### TRG Checklist for current state
#### TRG 1 Documentation

- [ ] [TRG 1.01](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-1) appropriate `README.md`
  - [ ] Basic description of repository and its content
  - [ ] Installation instructions to get component working
  - [ ] If required, additional post installation configuration steps to finish installation
- [ ] [TRG 1.02](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-2) appropriate install instructions either `INSTALL.md` or in `README.md`
  - [ ] Install instructions were extracted to into INSTALL.md 
- [ ] [TRG 1.03](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-3) appropriate `CHANGELOG.md`
- [x] [TRG 1.04](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-4) Diagrams as code / Editable static files 

#### TRG 2 Git

- [x] [TRG 2.01](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-1) default branch is named `main`
- [ ] [TRG 2.03](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-3) repository structure
  - [x] /docs
  - [ ] /charts
  - [ ] AUTHORS.md
  - [ ] CODE_OF_CONDUCT.md
  - [ ] CONTRIBUTING.md
  - [x] LICENSE
  - [ ] NOTICE.md
  - [ ] README.md
  - [ ] INSTALL.md
  - [ ] SECURITY.md
- [ ] [TRG 2.04](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-4) leading product repository
  - [x] README.md: contains the urls for the backend and frontend applications
  - [ ] contains the release of the product
  - [ ] contains the product helm chart
- [ ] [TRG 2.05](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-5) `.tractusx` metafile in a proper format
  - [ ] `.tractusx` file added
  - [ ] not published docker images listed

#### TRG 3 Kubernetes

- [ ] [TRG 3.02](https://eclipse-tractusx.github.io/docs/release/trg-3/trg-3-2) persistent volume and persistent volume claim is used when needed

#### TRG 4 Container

- [ ] [TRG 4.01](https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-01) [semantic versioning](https://semver.org/) and tagging <!-- container is tagged correctly additionally to the latest tag -->
- [ ] [TRG 4.02](https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-02) base image is agreed  <!-- Java, Kotlin, ... if JVM based language use base image from [Eclipse Temurin](https://hub.docker.com/_/eclipse-temurin) -->
- [ ] [TRG 4.03](https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-03) image has `USER` command and Non Root Container
- [ ] [TRG 4.05](https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-05) released image must be placed in `DockerHub`, remove `GHCR` references
- [ ] [TRG 4.06](https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-06) separate notice file for `DockerHub` has all necessary information

#### TRG 5 Helm

- [ ] [TRG 5.01](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-01) Helm chart must be released
- [ ] [TRG 5.02](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-02) Helm chart location in `/charts` directory and correct structure
- [ ] [TRG 5.04](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-04) CPU / MEM resource requests and limits and are properly set
- [ ] [TRG 5.06](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-06) Application must be configurable through the Helm chart <!-- every startup configuration aspect of your application must be configurable through the Helm chart (ingress class, tls, labels, annotations, database, secrets, persistence, env variables) -->
- [ ] [TRG 5.07](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-07) Dependencies are present and properly configured in the Chart.yaml
- [ ] [TRG 5.08](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-08) Product has a single deployable helm chart that contains all components <!--(backend, frontend, etc.) -->
- [ ] [TRG 5.09](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-09) Helm Test running properly
- [ ] [TRG 5.10](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-10) Products need to support 3 versions at a time
- [ ] [TRG 5.11](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-11) Upgradeability

#### TRG 6 Released Helm Chart

- [ ] [TRG 6.01](https://eclipse-tractusx.github.io/docs/release/trg-6/trg-6-1) Released Helm Chart <!-- A released Helm chart for each Tractus-X sub-product is expected to be available in corresponding GitHub repository. -->

#### TRG 7 Open Source Governance
- [ ] [TRG 7.01](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-01) Legal Documentation
- [ ] [TRG 7.02](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-02) License and copyright header <!-- must be present in every file if possible and update the year in the copyright section at the beginning of each new year. -->
- [ ] [TRG 7.03](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-03) IP checks for project content <!-- for each PR containing more than 1000 relevant lines there **must** be an approved [IP review for Code Contributions](/docs/oss/issues#eclipse-gitlab-ip-issue-tracker) before the contribution can be pushed/merged -->
- [ ] [TRG 7.04](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-04) IP checks for 3rd party content
- [ ] [TRG 7.05](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-05) Legal information for distributions
- [ ] [TRG 7.06](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-06) Legal information for end user content
- [ ] [TRG 7.07](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07) Legal notice for documentation

### Tasks to do before going to production
- [ ] **Mandatory!** Correctly release container images
- [ ] **Mandatory!** Correctly release helm chart
- [ ] **Mandatory!** Execute a long-term simulation test to see, if the feedback loop is stable.
- [ ] **Mandatory!** More elaborate testing of the `x/ubcmm` module
- [ ] **Mandatory!** Complete implementation of the `shiftup` operation at the `x/ubcmm` module.
- [ ] Create IaC (e.g. Terraform template) to allocate the necessary resources.
    - [ ] Sentry nodes layer to mitigate DDoS attacks
    - [ ] Use of TM-KMS to protect the validators block-signing-key and allow the use of hot fail-over validator without the risk of double signing
    - [ ] The sufficient allocation of hardware resources to the validator nodes.
- [ ] Distributed deployment of the different validator nodes (currently all at one Kubernetes cluster)
- [ ] Add a backup solution of the nodes' data directory.
- [ ] Setup a block-explorer for the network again.
- [ ] Reintegrate the cosmwasm module.


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
