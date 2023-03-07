// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

func NewResourceMap(resource Resource) ResourceMap {
	return ResourceMap{
		Resource:  resource,
		AuditLogs: nil,
	}
}
