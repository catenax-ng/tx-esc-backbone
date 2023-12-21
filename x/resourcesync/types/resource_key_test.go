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
	"github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

func TestNewResourceKey(t *testing.T) {
	type args struct {
		originator string
		origResId  string
	}
	tests := []struct {
		name    string
		args    args
		want    ResourceKey
		wantErr bool
	}{
		{
			"empty address fails",
			args{
				"",
				"some resId",
			},
			nil,
			true,
		},
		{
			"invalid address fails",
			args{
				"some invalid address",
				"some other resId",
			},
			nil,
			true,
		},
		{
			"Alice's address is fine",
			args{
				Alice,
				"some resId",
			},
			&resourceKey{
				Alice,
				"some resId",
			},
			false,
		},
		{
			"Bob's address is fine",
			args{
				Bob,
				"some other resId",
			},
			&resourceKey{
				Bob,
				"some other resId",
			},
			false,
		},
		{
			"Carol's address is fine",
			args{
				Carol,
				"yet another resId",
			},
			&resourceKey{
				Carol,
				"yet another resId",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourceKey(tt.args.originator, tt.args.origResId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResourceKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResourceKeyForDelete(t *testing.T) {
	type args struct {
		deleteMsg *MsgDeleteResource
	}
	tests := []struct {
		name    string
		args    args
		want    ResourceKey
		wantErr bool
	}{
		{
			"empty address fails",
			args{
				&MsgDeleteResource{
					"msgCreator",
					"",
					"some resId",
				},
			},
			nil,
			true,
		},
		{
			"invalid address fails",
			args{
				&MsgDeleteResource{
					"msgCreator",
					"some invalid address",
					"some other resId",
				},
			},
			nil,
			true,
		},
		{
			"Alice's address is fine",
			args{
				&MsgDeleteResource{
					"msgCreator",
					Alice,
					"some resId",
				},
			},
			&resourceKey{
				Alice,
				"some resId",
			},
			false,
		},
		{
			"Bob's address is fine",
			args{
				&MsgDeleteResource{
					"msgCreator",
					Bob,
					"some other resId",
				},
			},
			&resourceKey{
				Bob,
				"some other resId",
			},
			false,
		},
		{
			"Carol's address is fine",
			args{
				&MsgDeleteResource{
					"msgCreator",
					Carol,
					"yet another resId",
				},
			},
			&resourceKey{
				Carol,
				"yet another resId",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourceKeyForDelete(tt.args.deleteMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResourceKeyForDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceKeyForDelete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkOriginator(t *testing.T) {
	type args struct {
		originator string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty address fails",
			args{
				"",
			},
			true,
		},
		{
			"invalid address fails",
			args{
				"some invalid address",
			},
			true,
		},
		{
			"Alice's address is fine",
			args{
				Alice,
			},
			false,
		},
		{
			"Bob's address is fine",
			args{
				Bob,
			},
			false,
		},
		{
			"Carol's address is fine",
			args{
				Carol,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkOriginator(tt.args.originator); (err != nil) != tt.wantErr {
				t.Errorf("checkOriginator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ResourceKey_GetOrigResKey(t *testing.T) {
	type fields struct {
		originator string
		origResId  string
	}
	tests := []struct {
		name           string
		fields         fields
		wantOrigResKey string
	}{
		{
			"empty is fine",
			fields{
				Alice,
				"",
			},
			"",
		},
		{
			"some resId is fine",
			fields{
				Alice,
				"some resId",
			},
			"some resId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := resourceKey{
				originator: tt.fields.originator,
				origResId:  tt.fields.origResId,
			}
			if gotOrigResKey := m.GetOrigResKey(); gotOrigResKey != tt.wantOrigResKey {
				t.Errorf("GetOrigResKey() = %v, want %v", gotOrigResKey, tt.wantOrigResKey)
			}
		})
	}
}

func Test_ResourceKey_GetOriginator(t *testing.T) {
	type fields struct {
		originator string
		origResId  string
	}
	tests := []struct {
		name           string
		fields         fields
		wantOriginator string
	}{
		{
			"Alice's address is fine",
			fields{
				Alice,
				"some resId",
			},
			Alice,
		},
		{
			"Bob's address is fine",
			fields{
				Bob,
				"some other resId",
			},
			Bob,
		},
		{
			"Bob's address is fine",
			fields{
				Carol,
				"yet another resId",
			},
			Carol,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := resourceKey{
				originator: tt.fields.originator,
				origResId:  tt.fields.origResId,
			}
			if gotOriginator := m.GetOriginator(); gotOriginator != tt.wantOriginator {
				t.Errorf("GetOriginator() = %v, want %v", gotOriginator, tt.wantOriginator)
			}
		})
	}
}

func Test_ResourceKey_GetOriginatorAddress(t *testing.T) {
	aliceAccAddr, err := types.AccAddressFromBech32(Alice)
	if err != nil {
		panic(err)
	}
	bobAccAddr, err := types.AccAddressFromBech32(Bob)
	if err != nil {
		panic(err)
	}
	carolAccAddr, err := types.AccAddressFromBech32(Carol)
	if err != nil {
		panic(err)
	}
	type fields struct {
		originator string
		origResId  string
	}
	tests := []struct {
		name           string
		fields         fields
		wantOriginator types.AccAddress
	}{
		{
			"Alice's address is fine",
			fields{
				Alice,
				"some resId",
			},
			aliceAccAddr,
		},
		{
			"Bob's address is fine",
			fields{
				Bob,
				"some other resId",
			},
			bobAccAddr,
		},
		{
			"Carol's address is fine",
			fields{
				Carol,
				"some other resId",
			},
			carolAccAddr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := resourceKey{
				originator: tt.fields.originator,
				origResId:  tt.fields.origResId,
			}
			if gotOriginator := m.GetOriginatorAddress(); !reflect.DeepEqual(gotOriginator, tt.wantOriginator) {
				t.Errorf("GetOriginatorAddress() = %v, want %v", gotOriginator, tt.wantOriginator)
			}
		})
	}
}
