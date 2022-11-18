#!/bin/bash

# Copyright (c) 2022 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")


DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-node -t esc-backbone-node -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-node ghcr.io/catenax-ng/esc-backbone-node:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-node-init -t esc-backbone-node-init  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-node-init ghcr.io/catenax-ng/esc-backbone-node-init:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-web -t esc-backbone-web  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-web ghcr.io/catenax-ng/esc-backbone-web:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-faucet -t catena-faucet  -f "$SCRIPT_LOCATION/Dockerfile-faucet" "$SCRIPT_LOCATION/.."
docker tag catena-faucet ghcr.io/catenax-ng/catena-faucet:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-orchestrator -t esc-backbone-orchestrator  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-orchestrator ghcr.io/catenax-ng/esc-backbone-orchestrator:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  ssh-client -t ssh-client  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag ssh-client ghcr.io/catenax-ng/ssh-client:latest
