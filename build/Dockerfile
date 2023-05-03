# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0
# syntax=docker/dockerfile:1.4
FROM ghcr.io/catenax-ng/esc-backbone-build-base:latest as build

WORKDIR /build

RUN curl https://get.ignite.com/cli@v0.26.1! | bash

WORKDIR /build/esc-backbone
COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY app ./app
COPY docs ./docs
COPY x ./x
COPY testutil ./testutil
COPY test ./test
COPY proto ./proto
COPY config.yml ./

RUN ignite chain build
RUN go build github.com/catenax/esc-backbone/cmd/keplr-suggestion
COPY web ./web
RUN chmod u+x keplr-suggestion && mkdir -p ./web/public/chain && ./keplr-suggestion

FROM node:17.9 as web-chain-suggestion-build

WORKDIR /build
COPY --from=build /build/esc-backbone /build
WORKDIR /build/web
RUN npm install && npm run build

FROM nginxinc/nginx-unprivileged:1.23 as esc-backbone-web

ARG UID=101
ARG GID=101
USER 0
RUN apt-get update && apt-get install -y jq && apt-get remove --purge --auto-remove -y
ENV CHAIN_SUGGESTION=/usr/share/nginx/html/chain/catenax-testnet-1-suggestion.json
ENV CHAIN_REST="http://0.0.0.0:1317/"
ENV CHAIN_RPC="http://0.0.0.0:26657/"
ENV CHAIN_ID=""
ENV CHAIN_NAME=""
ENV CHAIN_DENOM_FILE=""
ENV CHAIN_DENOM_NEW=""
COPY --from=web-chain-suggestion-build /build/web/build /usr/share/nginx/html
RUN chown -R $UID:$GID /usr/share/nginx/html/chain/
COPY --chown=$UID:$GID build/webapp-entrypoint.sh /
COPY --chown=$UID:$GID build/write-env-to-json.sh /
COPY --chown=$UID:$GID build/update-chain-suggestion.sh  /

ENTRYPOINT ["/webapp-entrypoint.sh"]
USER $UID


CMD ["nginx", "-g", "daemon off;"]

FROM ghcr.io/catenax-ng/esc-backbone-debian-base:latest as esc-backbone-node-init
ARG BRANCH=main
ARG UID=102
ARG GID=102
LABEL branch="$BRANCH" repository="github.com/catenax-ng/product-esc-backbone-code"

ENV PACKAGES openssh-client sed jq git
RUN apt-get install -y $PACKAGES

# NODE_HOME specifies the folder name, which the validator will use as home folder below /validator/data
ENV NODE_HOME="chain"
RUN groupadd -g $GID validator
RUN useradd -g $GID -u $UID -s /bin/bash -d /home/validator validator
WORKDIR /initscripts
COPY --chown=$UID:$GID testnetscripts/*.sh /initscripts/
RUN chmod u+x /initscripts/*.sh
WORKDIR /validator
COPY --from=build --chown=$UID:$GID /output/dasel /usr/local/bin/
COPY --from=build --chown=$UID:$GID /go/bin/esc-backboned /validator/
COPY --from=build /build/esc-backbone/shared-libs /usr/lib
ENV CHAIN_BINARY="/validator/esc-backboned"
ENV REPO_BRANCH="main"
ENV CHAIN_ID="catenax-testnet-1"
ENV CURRENCY="ncaxdemo"
ENV MONIKER=""
# Seconds per try to wait for the git repo to be available
ENV GIT_WAIT=1
# Max retries on reading the git repo
ENV GIT_WAIT_MAX_RETRY=5
# Max retries for pushing a change to the sync repo (should be more than initial validators)
ENV GIT_PUSH_MAX_RETRY=5
# The public address the validator will be available at.
ENV PUBLIC_HOST_NAME=""
# The validator's mnemonic to create it's validator account key
ENV VALIDATOR_MNEMONIC=""
# The validator's mnemonic to create it's validator node key
ENV NODE_MNEMONIC=""

USER $UID

ENTRYPOINT ["sh", "-c","/initscripts/setup-node.sh /home/validator/$NODE_HOME $PUBLIC_HOST_NAME $MONIKER \"$VALIDATOR_MNEMONIC\" \"$NODE_MNEMONIC\""]

FROM ghcr.io/catenax-ng/esc-backbone-debian-base:latest as esc-backbone-node
ARG BRANCH=main
ARG UID=102
ARG GID=102

LABEL branch="$BRANCH" repository="github.com/catenax-ng/product-esc-backbone-code"
EXPOSE 1317 9090 9091 26656 26657 26658
# NODE_HOME specifies the folder name, which the validator will use as home folder below /home/validator
ENV NODE_HOME="chain"
RUN groupadd -g $GID validator
RUN useradd -g $GID -u $UID -s /bin/bash -d /home/validator validator
WORKDIR /validator
COPY --from=build --chown=$UID:$GID /go/bin/esc-backboned /validator/

USER $UID

ENTRYPOINT ["sh", "-c", "id; ls -la /home/validator; /validator/esc-backboned start --home /home/validator/$NODE_HOME"]

FROM ghcr.io/catenax-ng/esc-backbone-debian-base:latest as esc-backbone-orchestrator
ARG BRANCH=main
ARG UID=102
ARG GID=102
LABEL branch="$BRANCH" repository="github.com/catenax-ng/product-esc-backbone-code"

ENV PACKAGES  openssh-client git sed jq
RUN apt-get install -y $PACKAGES


# NODE_HOME specifies the folder name, which the validator will use as home folder below /validator/data
ENV NODE_HOME="chain"
RUN groupadd -g $GID validator
RUN useradd -g $GID -u $UID -s /bin/bash -d /home/validator validator
WORKDIR /initscripts
COPY --chown=$UID:$GID testnetscripts/*.sh /initscripts/
RUN chmod u+x /initscripts/*.sh
WORKDIR /orchestrator
RUN chown $UID:$GID /orchestrator
COPY --from=build --chown=$UID:$GID /output/dasel /usr/local/bin/
COPY --from=build --chown=$UID:$GID /go/bin/esc-backboned /orchestrator/
COPY --from=build /build/esc-backbone/shared-libs /usr/lib
ENV CHAIN_BINARY="/orchestrator/esc-backboned"
ENV REPO_BRANCH="main"
ENV CHAIN_ID="catenax-testnet-1"
ENV CURRENCY="ncaxdemo"
# Seconds per try to wait for the git repo to be available
ENV GIT_WAIT=1
# Max retries on reading the git repo
ENV GIT_WAIT_MAX_RETRY=5
# Max retries for pushing a change to the sync repo (should be more than initial validators)
ENV GIT_PUSH_MAX_RETRY=5

USER $UID
ENTRYPOINT ["sh", "-c","/initscripts/orchestrator-network-setup.sh /orchestrator/$NODE_HOME FALSE"]

FROM ghcr.io/catenax-ng/esc-backbone-debian-base:latest as ssh-client
ARG BRANCH=main

LABEL branch="$BRANCH" repository="github.com/catenax-ng/product-esc-backbone-code"

ENV PACKAGES  openssh-client sshpass curl jq git sed

RUN apt-get install -y $PACKAGES

COPY --from=build --chown=$UID:$GID /output/dasel /usr/local/bin/

# add the testnet scripts for easier watching a git repo
WORKDIR /testnetscripts
COPY testnetscripts/*.sh .
WORKDIR /
