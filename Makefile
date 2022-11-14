.PHONY: build-vmso

VER_TAG=v1.0.0

build-vmso:
	git clone https://github.com/CosmWasm/wasmvm wasmvm
	(cd wasmvm && git fetch --all --tags && git checkout $(VER_TAG))
	(cd wasmvm && make build-rust)