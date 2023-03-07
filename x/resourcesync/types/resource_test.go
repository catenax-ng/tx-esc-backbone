// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

func TestResource_GetDataHash(t *testing.T) {
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name         string
		fields       fields
		wantDataHash []byte
		wantErr      bool
	}{
		{
			"nil is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				nil,
			},
			nil,
			false,
		},
		{
			"empty is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte{},
			},
			[]byte{},
			false,
		},
		{
			"some short hash is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte{5, 5, 5},
			},
			[]byte{5, 5, 5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			gotDataHash, err := m.GetDataHash()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDataHash, tt.wantDataHash) {
				t.Errorf("GetDataHash() gotDataHash = %v, want %v", gotDataHash, tt.wantDataHash)
			}
		})
	}
}

func TestResource_GetOrigResKey(t *testing.T) {
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name           string
		fields         fields
		wantOrigResKey string
		wantErr        bool
	}{
		{
			"empty is fine",
			fields{
				Alice,
				"",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			"",
			false,
		},
		{
			"some resId is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			"some resId",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			gotOrigResKey, err := m.GetOrigResKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrigResKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOrigResKey != tt.wantOrigResKey {
				t.Errorf("GetOrigResKey() gotOrigResKey = %v, want %v", gotOrigResKey, tt.wantOrigResKey)
			}
		})
	}
}

func TestResource_GetOriginatorAddress(t *testing.T) {
	aliceAccAddr, err := types.AccAddressFromBech32(Alice)
	if err != nil {
		panic(err)
	}
	bobAccAddr, err := types.AccAddressFromBech32(Bob)
	if err != nil {
		panic(err)
	}
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name           string
		fields         fields
		wantOriginator types.AccAddress
		wantErr        bool
	}{
		{
			"empty address fails",
			fields{
				"",
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			types.AccAddress{},
			true,
		},
		{
			"invalid address fails",
			fields{
				"some invalid address",
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			nil,
			true,
		},
		{
			"Alice's address is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			aliceAccAddr,
			false,
		},
		{
			"Bob's address is fine",
			fields{
				Bob,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			bobAccAddr,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			gotOriginator, err := m.GetOriginatorAddress()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginatorAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOriginator, tt.wantOriginator) {
				t.Errorf("GetOriginatorAddress() gotOriginator = %v, want %v", gotOriginator, tt.wantOriginator)
			}
		})
	}
}

func TestResource_GetResourceKey(t *testing.T) {
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name            string
		fields          fields
		wantResourceKey string
		wantErr         bool
	}{
		{
			"empty is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"",
				[]byte("some hash"),
			},
			"",
			false,
		},
		{
			"some resource key is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			"some resource key",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			gotResourceKey, err := m.GetResourceKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResourceKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResourceKey != tt.wantResourceKey {
				t.Errorf("GetResourceKey() gotResourceKey = %v, want %v", gotResourceKey, tt.wantResourceKey)
			}
		})
	}
}

func TestResource_GetTargetSystem(t *testing.T) {
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name             string
		fields           fields
		wantTargetSystem string
		wantErr          bool
	}{
		{
			"empty is fine",
			fields{
				Alice,
				"some resId",
				"",
				"some resource key",
				[]byte("some hash"),
			},
			"",
			false,
		},
		{
			"some target system is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			"some target system",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			gotTargetSystem, err := m.GetTargetSystem()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTargetSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTargetSystem != tt.wantTargetSystem {
				t.Errorf("GetTargetSystem() gotTargetSystem = %v, want %v", gotTargetSystem, tt.wantTargetSystem)
			}
		})
	}
}

func TestResource_ToResourceKey(t *testing.T) {
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name            string
		fields          fields
		wantResourceKey ResourceKey
		wantErr         bool
	}{
		{
			"empty address fails",
			fields{
				"",
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			nil,
			true,
		},
		{
			"invalid address fails",
			fields{
				"some invalid address",
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			nil,
			true,
		},
		{
			"Alice's address is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			&resourceKey{
				Alice,
				"some resId",
			},
			false,
		},
		{
			"Bob's address is fine",
			fields{
				Bob,
				"some other resId",
				"this value has no impact",
				"this value has no effect",
				[]byte("some hash"),
			},
			&resourceKey{
				Bob,
				"some other resId",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			gotResourceKey, err := m.ToResourceKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToResourceKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResourceKey, tt.wantResourceKey) {
				t.Errorf("ToResourceKey() gotOriginator = %v, want %v", gotResourceKey, tt.wantResourceKey)
			}
		})
	}
}

func TestResource_Validate(t *testing.T) {
	type fields struct {
		Originator   string
		OrigResId    string
		TargetSystem string
		ResourceKey  string
		DataHash     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"some resource is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			false,
		},
		{
			"with no address fails",
			fields{
				"",
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			true,
		},
		{
			"with empty origResId is fine",
			fields{
				Alice,
				"",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			false,
		},
		{
			"with empty targetSystem is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			false,
		}, {
			"with empty resourceKey is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte("some hash"),
			},
			false,
		},
		{
			"with empty dataHash is fine",
			fields{
				Alice,
				"some resId",
				"some target system",
				"some resource key",
				[]byte{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Resource{
				Originator:   tt.fields.Originator,
				OrigResId:    tt.fields.OrigResId,
				TargetSystem: tt.fields.TargetSystem,
				ResourceKey:  tt.fields.ResourceKey,
				DataHash:     tt.fields.DataHash,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
