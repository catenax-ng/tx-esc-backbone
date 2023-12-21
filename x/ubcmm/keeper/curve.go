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
package keeper

import (
	"github.com/catenax/esc-backbone/x/ubcmm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCurve set curve in the store
func (k Keeper) SetCurve(ctx sdk.Context, curve types.Curve) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurveKey))
	b := k.cdc.MustMarshal(&curve)
	store.Set([]byte{0}, b)
}

// GetCurve returns curve
func (k Keeper) GetCurve(ctx sdk.Context) (val types.Curve, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurveKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	val.PopulateSegments()
	return val, true
}

// RemoveCurve removes curve from the store
func (k Keeper) RemoveCurve(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurveKey))
	store.Delete([]byte{0})
}
