// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ResourceMapList: []ResourceMap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in resourceMap
	resourceMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.ResourceMapList {
		index := string(ResourceMapKey(elem.Originator, elem.OrigResId))
		if _, ok := resourceMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for resourceMap")
		}
		resourceMapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
