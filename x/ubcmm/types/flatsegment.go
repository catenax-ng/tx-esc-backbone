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

import sdk "github.com/cosmos/cosmos-sdk/types"

// setP0X is a noop, as the P0 is always 0,0 and hence not stored.
func (fseg *FlatSegment) setP0X(_ sdk.Dec) {
}

// setP0Y is a noop, as the P0 is always 0,0 and hence not stored.
func (fseg *FlatSegment) setP0Y(_ sdk.Dec) {
}

// setP1X sets the value of parameter p1Y.
func (fseg *FlatSegment) setP1Y(p1Y sdk.Dec) {
	fseg.Y = p1Y
}

// setP1X sets the value of parameter p1X.
func (fseg *FlatSegment) setP1X(p1X sdk.Dec) {
	fseg.P1X = p1X
}

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2.
func (fseg *FlatSegment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	return fseg.Y.Mul(x2.Sub(x1))
}

// firstDerivativeX1 returns 0, since the segment is a horizontal line.
func (fseg *FlatSegment) firstDerivativeX1(_ sdk.Dec) (_ sdk.Dec) {
	return sdk.ZeroDec()
}

// y returns the y value for the given x.
func (fseg *FlatSegment) y(x sdk.Dec) sdk.Dec {
	return fseg.Y
}

var _ view = (*FlatSegment)(nil)

// startX returns the x-value of the start of the visible part of the line
func (fseg *FlatSegment) startX() sdk.Dec {
	return sdk.ZeroDec() // TODO: Check if this is correct.
}

// endX returns the x-value of the end of the visible part of the line.
func (fseg *FlatSegment) endX() sdk.Dec {
	return fseg.P1X
}

// startY returns the y-value of the start of the visible part of the line
func (fseg *FlatSegment) startY() sdk.Dec {
	return fseg.Y
}

// endY returns the y-value of the end of the visible part of the line.
func (fseg *FlatSegment) endY() sdk.Dec {
	return fseg.Y
}
