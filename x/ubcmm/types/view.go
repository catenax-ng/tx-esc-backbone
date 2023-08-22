// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// view provides methods for accessing the visible part of a curve.
//
// Note: For curves with
// - dynamic shape: curve is visible across its entire x-interval
// - fixed shape: curve is visible only across a part of its x-interval.
type view interface {
	startX() sdk.Dec
	endX() sdk.Dec

	startY() sdk.Dec
	endY() sdk.Dec
}
