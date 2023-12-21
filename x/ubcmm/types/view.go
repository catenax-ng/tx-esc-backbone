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
