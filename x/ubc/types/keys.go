// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

const (
	// ModuleName defines the module name
	ModuleName = "ubc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ubc"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	UbcobjectKey = "Ubcobject/value/"
)
