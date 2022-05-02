FROM golang:1.17.8-buster as build

ENV PACKAGES git openssh-client bash curl

WORKDIR /build
RUN apt update
RUN apt install $PACKAGES
RUN curl https://get.starport.network/starport | bash && chown root:root starport
RUN mkdir -p -m 0700 ~/.ssh \
    #    ssh-keyscan github.com >> ~/.ssh/known_hosts # Did not work. \
    && echo "|1|XGM2kbi/MRZnS54PMAWDoGNMpy4=|ZA+nLvP/UQ19SnwZjlnP8rjY364= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=" >> ~/.ssh/known_hosts


ARG BRANCH=main
RUN --mount=type=secret,id=ssh_key ssh-agent ssh-agent sh -c 'ssh-add /run/secrets/ssh_key && git clone -b $BRANCH git@github.com:catenax/esc-backbone.git'
#RUN --mount=type=ssh git clone git@github.com:catenax/esc-backbone.git

WORKDIR /build/esc-backbone

RUN /build/starport chain build -v

RUN which esc-backboned
#RUN /build/starport chain build --help

FROM debian:10.11-slim as faucet
# TODO implement command for the faucet

FROM debian:10.11-slim as vue
WORKDIR /vue
COPY --from=build /build/esc-backbone/vue /vue
# TODO serve the vue app

ENTRYPOINT ["npm","run","serve"]

FROM debian:10.11-slim as faucet
# TODO implement command for the faucet

FROM debian:10.11-slim as blockexpl

FROM debian:10.11-slim as node
ARG BRANCH=main
LABEL branch="$BRANCH" repository="github.com/catenax/esc-backbone"
EXPOSE 1317 26656 26657 26658
RUN useradd --help
RUN useradd -U -s /bin/bash -d /validator validator
WORKDIR /validator
COPY --from=build /go/bin/esc-backboned /validator/
VOLUME /validator/chain/
RUN chown -R validator:validator /validator
USER validator
# TODO REQUEST funds from faucet
# TODO specify validator key using the /run/secret/mnemonic file
ENTRYPOINT ["/validator/esc-backboned","start","--home","/validator/chain"]