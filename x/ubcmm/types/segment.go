// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Segment interface {
		setP0X(sdk.Dec)
		setP1X(sdk.Dec)
		setP0Y(sdk.Dec)
		setP1Y(sdk.Dec)

		y(x sdk.Dec) sdk.Dec
		integralX12(x1, x2 sdk.Dec) sdk.Dec
		firstDerivativeX1(x1 sdk.Dec) sdk.Dec

		view
	}

	// Segments is a special type for defining a gogo proto compatible field in the Curve to store a list of Segment.
	//
	// This types always marshals to zero, since it is not intended to be stored on the blockchain.
	//
	// On unmarshalling, it is set to nil. It should be populated using the segment data stored in the curve.
	Segments []Segment
)

// MarshalTo implements the gogo proto custom type interface.
//
// Its a no-op always returns (0, nil), since this type is not intended to be
// stored on the blockchain.
func (_ Segments) MarshalTo(data []byte) (n int, err error) {
	return 0, nil
}

// Unmarshal implements the gogo proto custom type interface.
//
// Its a no-op and always returns nil, since this type is not intended to be
// stored on the blockchain.
func (_ Segments) Unmarshal(_ []byte) error {
	return nil
}

// Size implements the gogo proto custom type interface.
//
// Its a no-op and always returns 0, since this type is not intended to be
// stored on the blockchain.
func (_ Segments) Size() int {
	return 0
}
