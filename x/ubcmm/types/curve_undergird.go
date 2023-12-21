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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (c *Curve) UndergirdS02(BPoolAdd sdk.Dec) error {
	if c.CurrentSupply.LT(c.pX(2)) {
		errMsg := "could not undergird, since the currentSupply is not beyond P2"
		return sdkerrors.ErrInvalidRequest.Wrap(errMsg)
	}

	c.BPoolUnder = c.BPoolUnder.Add(BPoolAdd)
	c.BPool = c.BPool.Add(BPoolAdd)

	return c.FitUntilConvergence()
}
