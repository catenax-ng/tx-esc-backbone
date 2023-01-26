package types

import (
	"reflect"
	"testing"
)

func TestResourceMapKey(t *testing.T) {
	type args struct {
		originator string
		origResId  string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"originator contributes to key",
			args{
				"orig1",
				"",
			},
			[]byte("orig1//"),
		},
		{
			"origResId contributes to key",
			args{
				"",
				"orig1key",
			},
			[]byte("/orig1key/"),
		},
		{
			"originator and origResId contributes to key",
			args{
				"orig2",
				"orig2key",
			},
			[]byte("orig2/orig2key/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResourceMapKey(tt.args.originator, tt.args.origResId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceMapKeyOf(t *testing.T) {
	reusedResourceKey := createValidResouceKey(Alice, "reusedOrigKey")

	type args struct {
		resource ResourceKey
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"originator and origResId contribute to key",
			args{
				createValidResouceKey(Bob, "orig2key"),
			},
			[]byte(Bob + "/orig2key/"),
			false,
		},
		{
			"only originator and origResId contribute to key other fields are ignored",
			args{
				createValidResouceKey(Bob, "orig2key"),
			},
			[]byte(Bob + "/orig2key/"),
			false,
		},
		{
			"Delivers the same as ResourceMapKey(reusedResourceKey.GetOriginator(),reusedResourceKey.GetOrigResKey())",
			args{
				reusedResourceKey,
			},
			ResourceMapKey(reusedResourceKey.GetOriginator(), reusedResourceKey.GetOrigResKey()),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResourceMapKeyOf(tt.args.resource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceMapKeyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
