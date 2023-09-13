// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Curve_Undergird(t *testing.T) {
	c := validCurve()
	require.NoError(t, c.Fit())
	require.NoError(t, c.UndergirdS01(sdk.NewDec(100e5)))
	assert.True(t, c.IsIntegralEqualToBPool())
}
