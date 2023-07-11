#!/bin/bash
# Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
#
# See the NOTICE file(s) distributed with this work for additional
# information regarding copyright ownership.
#
# This program and the accompanying materials are made available under the
# terms of the Apache License, Version 2.0 which is available at
# https://www.apache.org/licenses/LICENSE-2.0.
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.
#
# SPDX-License-Identifier: Apache-2.0


SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")

DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-build-base -t esc-backbone-build-base -f "$SCRIPT_LOCATION/Dockerfile-base" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-build-base ghcr.io/catenax-ng/esc-backbone-build-base
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-debian-base -t esc-backbone-debian-base -f "$SCRIPT_LOCATION/Dockerfile-base" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-debian-base ghcr.io/catenax-ng/esc-backbone-debian-base
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-web-build-base -t esc-backbone-web-build-base -f "$SCRIPT_LOCATION/Dockerfile-base" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-web-build-base ghcr.io/catenax-ng/esc-backbone-web-build-base
DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-backbone-web-base -t esc-backbone-web-base -f "$SCRIPT_LOCATION/Dockerfile-base" "$SCRIPT_LOCATION/.."
docker tag esc-backbone-web-base ghcr.io/catenax-ng/esc-backbone-web-base

DOCKER_BUILDKIT=1 docker build --progress=plain --target  esc-res-sync-rest-wrapper -t esc-res-sync-rest-wrapper -f "$SCRIPT_LOCATION/Dockerfile" "$SCRIPT_LOCATION/.."
docker tag esc-res-sync-rest-wrapper ghcr.io/catenax-ng/esc-res-sync-rest-wrapper:latest

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
DOCKER_BUILDKIT=1 docker build --progress=plain --target  demo -t mid-term-demo  -f "${SCRIPT_LOCATION}/../web-mid-term-demo/build/Dockerfile" "${SCRIPT_LOCATION}/../web-mid-term-demo"
docker tag mid-term-demo ghcr.io/catenax-ng/esc-mid-term-demo:latest


