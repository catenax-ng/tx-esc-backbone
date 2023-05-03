#!/bin/bash
# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
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