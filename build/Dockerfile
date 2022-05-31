# syntax=docker/dockerfile:1.4
FROM golang:1.18.1-bullseye as build

ENV PACKAGES openssh-client bash curl

WORKDIR /build

RUN apt update
RUN apt install $PACKAGES

RUN curl https://get.ignite.com/cli?type=script | bash && chown root:root ignite

WORKDIR /build/esc-backbone

COPY cmd ./cmd
COPY app ./app
COPY web ./web
COPY docs ./docs
COPY go.mod go.sum config.yml ./

RUN /build/ignite chain build
RUN go build github.com/catenax/esc-backbone/cmd/keplr-suggestion
RUN chmod u+x keplr-suggestion && mkdir -p ./web/public/chain && ./keplr-suggestion

# COPY prepared test net homes.
COPY testnet /home/testnet

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
RUN chown $UID:$GID /usr/share/nginx/html/chain/*.json
COPY --chown=$UID:$GID build/webapp-entrypoint.sh /
COPY --chown=$UID:$GID build/update-chain-suggestion.sh  /

ENTRYPOINT ["/webapp-entrypoint.sh"]
USER $UID


CMD ["nginx", "-g", "daemon off;"]

FROM debian:10.11-slim as faucet
# TODO implement command for the faucet

FROM debian:10.11-slim as blockexpl

FROM debian:10.11-slim as esc-backbone-node
ARG BRANCH=main
ARG UID=102
ARG GID=102

LABEL branch="$BRANCH" repository="github.com/catenax/esc-backbone"
EXPOSE 1317 9090 9091 26656 26657 26658
# NODE_HOME specifies the folder name, which the validator will use as home folder below /validator/data
ENV NODE_HOME="chain"
# TEMPLATE_HOME can point to a home folder for the validator.
# The folder is copied into /validator/data/$NODE_HOME, if /validator/data/$NODE_HOME does not exist.
ENV TEMPLATE_HOME=""
RUN groupadd -g $GID validator
RUN useradd -g $GID -u $UID -s /bin/bash -d /validator validator
WORKDIR /validator
COPY --chown=$UID:$GID build/validator-node-entrypoint.sh /validator/
RUN chmod u+x /validator/validator-node-entrypoint.sh
COPY --from=build --chown=$UID:$GID /go/bin/esc-backboned /validator/
COPY --from=build --chown=$UID:$GID /home/testnet /home/testnet

VOLUME /validator/data/
USER $UID
# TODO REQUEST funds from faucet
# TODO specify validator key using the /run/secret/mnemonic file
ENTRYPOINT ["sh", "-c","/validator/validator-node-entrypoint.sh /validator/data $NODE_HOME $TEMPLATE_HOME"]
