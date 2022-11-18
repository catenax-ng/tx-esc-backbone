#!/bin/bash

# Copyright (c) 2022 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0

kind load docker-image ghcr.io/catenax-ng/esc-backbone-node:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-web:latest
kind load docker-image ghcr.io/catenax-ng/catena-faucet:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-node-init:latest
kind load docker-image ghcr.io/catenax-ng/esc-backbone-orchestrator:latest