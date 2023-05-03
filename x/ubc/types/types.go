// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

const (
	// CaxMultiplier is the multiplier for cax tokens
	CaxMultiplier = 1e18

	// VoucherMultiplier is the multiplier for voucher
	VoucherMultiplier = 1e2

	// CaxDenom is the denominator for cax tokens
	CaxDenom = "acax"

	// VoucherDenom is the denominator for voucher
	VoucherDenom = "cvoucher"
)

func isValidDenom(denom string) bool {
	return denom == VoucherDenom || denom == CaxDenom
}
