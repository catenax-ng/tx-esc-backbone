#!/bin/bash

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")


DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-node -t esc-backbone-node -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-node ghcr.io/catenax-ng/esc-backbone-node:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-web -t esc-backbone-web  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-web ghcr.io/catenax-ng/esc-backbone-web:latest
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-faucet -t catena-faucet  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-faucet ghcr.io/catenax-ng/catena-faucet:latest
