# Copyright (c) 2022-2023 - for information on the respective copyright owner
# see the NOTICE file and/or the repository at
# https://github.com/catenax-ng/product-esc-backbone-code
#
# SPDX-License-Identifier: Apache-2.0
.PHONY: build-vmso

VER_TAG=v1.0.0

build-vmso:
	git clone https://github.com/CosmWasm/wasmvm wasmvm
	(cd wasmvm && git fetch --all --tags && git checkout $(VER_TAG))
	(cd wasmvm && make build-rust)

gen-proto-go:
	ignite generate proto-go --yes
	ignite generate openapi --yes
	./license-header/update-year.sh
