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

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ResourceMapKeyPrefix is the prefix to retrieve all ResourceMap
	ResourceMapKeyPrefix = "ResourceMap/value/"
)

// ResourceMapKey returns the store key to retrieve a ResourceMap from the index fields
func ResourceMapKey(
	originator string,
	origResId string,
) []byte {
	var key []byte

	originatorBytes := []byte(originator)
	key = append(key, originatorBytes...)
	key = append(key, []byte("/")...)

	origResIdBytes := []byte(origResId)
	key = append(key, origResIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ResourceMapKeyOf(resource ResourceKey) []byte {
	return ResourceMapKey(resource.GetOriginator(), resource.GetOrigResKey())
}
