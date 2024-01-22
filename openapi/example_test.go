package openapi

import (
	"reflect"
	"testing"
)

func Test_newExample(t *testing.T) {
	type args struct {
		t Type
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{TypeString, "abc"},
			`"abc"`,
		},
		{
			args{TypeInteger, "123"},
			`123`,
		},
		{
			args{TypeNumber, "123.4"},
			`123.4`,
		},
		{
			args{TypeBoolean, "false"},
			`false`,
		},
		{
			args{TypeBoolean, "any string"},
			`any string`,
		},
		{
			args{TypeArray, "any string"},
			`any string`,
		},
		{
			args{TypeObject, "any string"},
			`any string`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			ex := newExample(tt.args.t, tt.args.s)
			want := []byte(tt.want)
			if !reflect.DeepEqual(ex.value, want) {
				t.Errorf("newExample() = %v, want %v", ex.value, want)
			}
		})
	}
}
