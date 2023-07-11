// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package main

// RequestResource is used for create, update and delete resources.
//
// swagger:model RequestResource
type RequestResource struct {
	// OrigResId is the id of the resource by originator -  unique per originator
	//
	// required: true
	OrigResId string `json:"origResId,omitempty"`

	// TargetSystem is holding the information of the resource
	//
	// required: true
	TargetSystem string `json:"targetSystem,omitempty"`
	// ResourceKey is the Id of the resource to access it at the target system
	//
	// required: true
	ResourceKey string `json:"resourceKey,omitempty"`
	// DataHash contains the hash of the resource for integrity check
	//
	// required: false
	DataHash []byte `json:"dataHash,omitempty"`
}
