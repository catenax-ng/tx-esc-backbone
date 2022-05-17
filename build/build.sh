#!/bin/bash

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")


DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-node -t esc-backbone -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-web -t esc-web  -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."