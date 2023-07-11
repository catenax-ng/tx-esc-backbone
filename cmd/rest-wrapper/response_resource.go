// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package main

// ResponseResource contains information about the issued transaction.
//
// swagger:model ResponseResource
type ResponseResource struct {
	// Height of the block containing the transaction
	Height int64 `json:"height,omitempty"`
	// TxHash is the hash of the transaction
	TxHash string `json:"txHash,omitempty"`
	// RawLog contains the raw log of the transaction.
	RawLog string `json:"rawLog,omitempty"`
}
