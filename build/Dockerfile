FROM golang:1.18.1-bullseye as build

ENV PACKAGES openssh-client bash curl

WORKDIR /build

RUN apt update
RUN apt install $PACKAGES
RUN curl https://get.ignite.com/cli | bash && chown root:root ignite

WORKDIR /build/esc-backbone

COPY cmd ./cmd
COPY app ./app
COPY web ./web
COPY docs ./docs
COPY go.mod go.sum config.yml ./

RUN /build/ignite chain build
RUN go build github.com/catenax/esc-backbone/cmd/keplr-suggestion
RUN chmod u+x keplr-suggestion && mkdir -p ./web/src/
# COPY prepared test net homes.
COPY testnet /home/testnet

FROM node:17.9 as web-chain-suggestion-build

WORKDIR /build
COPY --from=build /build/esc-backbone /build
WORKDIR /build/web
RUN npm install && npm run build

FROM nginx:1.21 as esc-backbone-web

COPY --from=web-chain-suggestion-build /build/web/dist /usr/share/nginx/html

FROM debian:10.11-slim as faucet
# TODO implement command for the faucet

FROM debian:10.11-slim as blockexpl

FROM debian:10.11-slim as esc-backbone-node
ARG BRANCH=main
LABEL branch="$BRANCH" repository="github.com/catenax/esc-backbone"
EXPOSE 1317 9090 9091 26656 26657 26658
ENV NODE_HOME=/validator/chain
RUN useradd -U -s /bin/bash -d /validator validator
WORKDIR /validator
COPY --from=build /go/bin/esc-backboned /validator/
RUN chown -R validator:validator /validator
COPY --from=build /home/testnet /home/testnet
RUN chown -R validator:validator /home/testnet

VOLUME /validator/chain/
USER validator
# TODO REQUEST funds from faucet
# TODO specify validator key using the /run/secret/mnemonic file
ENTRYPOINT ["sh", "-c","id; /validator/esc-backboned start --home $NODE_HOME"]