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

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ubc module sentinel errors
var (
	ErrSample       = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidArg   = sdkerrors.Register(ModuleName, 1101, "invalid argument")
	ErrCurveFitting = sdkerrors.Register(ModuleName, 1102, "fitting the curve")
	ErrFundHandling = sdkerrors.Register(ModuleName, 1103, "fund handling")
	ErrComputation  = sdkerrors.Register(ModuleName, 1104, "computation error")
)
