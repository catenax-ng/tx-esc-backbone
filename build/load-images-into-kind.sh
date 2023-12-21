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


kind load docker-image ghcr.io/catenax-ng/esc-backbone-build-base:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-debian-base:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-web-build-base:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-web-base:latest

kind load docker-image ghcr.io/catenax-ng/esc-backbone-node:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-web:latest
kind load docker-image ghcr.io/catenax-ng/catena-faucet:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-node-init:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-orchestrator:latest
kind load docker-image ghcr.io/catenax-ng/ssh-client:latest
kind load docker-image ghcr.io/catenax-ng/esc-mid-term-demo:latest