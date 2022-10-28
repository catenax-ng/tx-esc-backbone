# syntax=docker/dockerfile:1.4
FROM golang:1.19.1-bullseye as build

ENV PACKAGES bash curl

WORKDIR /build

RUN apt update
RUN apt install $PACKAGES -y

RUN curl https://get.ignite.com/cli@v0.26.1?type=script | bash && chown root:root ignite

WORKDIR /build/esc-backbone

COPY cmd ./cmd
COPY app ./app
COPY docs ./docs
COPY go.mod go.sum config.yml ./

RUN /build/ignite chain build
RUN go build github.com/catenax/esc-backbone/cmd/keplr-suggestion
COPY web ./web
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
RUN chown -R $UID:$GID /usr/share/nginx/html/chain/
COPY --chown=$UID:$GID build/webapp-entrypoint.sh /
COPY --chown=$UID:$GID build/write-env-to-json.sh /
COPY --chown=$UID:$GID build/update-chain-suggestion.sh  /

ENTRYPOINT ["/webapp-entrypoint.sh"]
USER $UID


CMD ["nginx", "-g", "daemon off;"]


FROM debian:10.11-slim as cosmjs-checkout
RUN apt update
RUN apt upgrade -y
RUN apt install -y git
WORKDIR /git
RUN git clone --branch v0.28.13 https://github.com/cosmos/cosmjs.git
# this is required to fix the build issues
RUN rm cosmjs/yarn.lock

# based on https://github.com/cosmos/cosmjs/blob/v0.28.13/packages/faucet/Dockerfile#L11-L20
FROM node:14-alpine as faucet-build-env
RUN apk add --update --no-cache alpine-sdk linux-headers build-base gcc libusb-dev python3 py3-pip eudev-dev
RUN ln -sf python3 /usr/bin/python
WORKDIR /app
COPY --from=cosmjs-checkout /git/cosmjs /app

RUN yarn install && yarn run build
RUN (cd packages/faucet && SKIP_BUILD=1 yarn pack-node)

# based on https://github.com/cosmos/cosmjs/blob/v0.28.13/packages/faucet/Dockerfile#L22-L32
FROM alpine:3.15 as esc-backbone-faucet
# Installs Node.js 16 (https://pkgs.alpinelinux.org/packages?name=nodejs&branch=v3.15)
RUN apk add --update nodejs

COPY --from=faucet-build-env /app/packages/faucet /app/packages/faucet
WORKDIR /app

EXPOSE 8000
ENV FAUCET_ADDRESS_PREFIX="catenax"
ENV FAUCET_TOKENS="ncaxdemo"
ENTRYPOINT ["/app/packages/faucet/bin/cosmos-faucet-dist"]

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
