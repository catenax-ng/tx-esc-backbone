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
	reusedResource := Resource{
		"reusedOrig",
		"reusedOrigKey",
		"reused ignored string",
		"reused ignored key",
		[]byte{42, 5, 7},
	}
	type args struct {
		resource *Resource
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"originator contributes to key",
			args{
				&Resource{
					"orig1",
					"",
					"",
					"",
					[]byte{},
				},
			},
			[]byte("orig1//"),
		},
		{
			"origResId contributes to key",
			args{
				&Resource{
					"",
					"orig1key",
					"",
					"",
					[]byte{},
				},
			},
			[]byte("/orig1key/"),
		},
		{
			"originator and origResId contribute to key",
			args{
				&Resource{
					"orig2",
					"orig2key",
					"",
					"",
					[]byte{},
				},
			},
			[]byte("orig2/orig2key/"),
		},
		{
			"only originator and origResId contribute to key other fields are ignored",
			args{
				&Resource{
					"orig2",
					"orig2key",
					"ignored string",
					"ignored key",
					[]byte{5, 7},
				},
			},
			[]byte("orig2/orig2key/"),
		}, {
			"Delivers the same as ResourceMapKey(resource.Originator,resource.OrigResId)",
			args{
				&reusedResource,
			},
			ResourceMapKey(reusedResource.Originator, reusedResource.OrigResId),
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
