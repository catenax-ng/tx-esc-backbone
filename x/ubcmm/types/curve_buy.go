// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Buy buys the given amount of tokens against the curve. It
// returns the amount of vouchers used.
//
// It assumes the value of tokens is greater than zero. This condition is
// implemented in the ValidBasic function for the buy message.
func (c *Curve) Buy(tokens sdk.Dec) sdk.Dec {
	xCurrent := c.CurrentSupply
	xNew := c.CurrentSupply.Add(tokens)

	vouchersUsed := c.integralX12(xCurrent, xNew)
	vouchersUsed = c.ubcRoundOff(vouchersUsed, VoucherMultiplier)

	c.CurrentSupply = xNew
	c.BPool = c.BPool.Add(vouchersUsed)
	// CLARIFY: Should we change BPoolUnder
	return vouchersUsed
}

func (c *Curve) ubcRoundOff(n sdk.Dec, multiplier int64) sdk.Dec {
	rounded := bankersRoundOff(n, multiplier)
	difference := n.Sub(rounded)
	c.NumericalErrorAccumulator = c.NumericalErrorAccumulator.Add(difference)
	return rounded
}

func bankersRoundOff(n sdk.Dec, multiplier int64) sdk.Dec {
	integer := n.MulInt64(multiplier)
	roundedInteger := integer.RoundInt()
	return sdk.NewDecFromInt(roundedInteger).QuoInt64(multiplier)
}
