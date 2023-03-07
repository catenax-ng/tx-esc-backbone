// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import "github.com/catenax/esc-backbone/testutil"

const (
	Alice = testutil.Alice
	Bob   = testutil.Bob
	Carol = testutil.Carol
)

func createValidResouceKey(originator string, origResId string) ResourceKey {
	resourceKey, err := NewResourceKey(originator, origResId)
	if err != nil {
		panic(err)
	}
	return resourceKey
}
