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
package web2wrapper

// RequestResource is used for create, update and delete resources.
//
// swagger:model RequestResource
type RequestResource struct {
	// OrigResId is the id of the resource by originator -  unique per originator
	//
	// required: true
	OrigResId string `json:"origResId,omitempty"`

	// TargetSystem is holding the information of the resource
	//
	// required: true
	TargetSystem string `json:"targetSystem,omitempty"`
	// ResourceKey is the Id of the resource to access it at the target system
	//
	// required: true
	ResourceKey string `json:"resourceKey,omitempty"`
	// DataHash contains the hash of the resource for integrity check
	//
	// required: false
	DataHash []byte `json:"dataHash,omitempty"`
}
