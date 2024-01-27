package schema

import (
	"reflect"
	"testing"
)

func Test_newStringExample(t *testing.T) {
	type args struct {
		astExampleString string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"abc"},
			`"abc"`,
		},
		{
			args{"123"},
			`"123"`,
		},
		{
			args{"123.4"},
			`"123.4"`,
		},
		{
			args{"false"},
			`"false"`,
		},
		{
			args{"any string"},
			`"any string"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			ex := newStringExample(tt.args.astExampleString)
			want := []byte(tt.want)
			if !reflect.DeepEqual(ex.value, want) {
				t.Errorf("newStringExample() = %v, want %v", ex.value, want)
			}
		})
	}
}

func Test_newBasicExample(t *testing.T) {
	type args struct {
		astExampleString string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"abc"},
			`abc`,
		},
		{
			args{"123"},
			`123`,
		},
		{
			args{"123.4"},
			`123.4`,
		},
		{
			args{"false"},
			`false`,
		},
		{
			args{"any string"},
			`any string`,
		},
		{
			args{"any string"},
			`any string`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			ex := newBasicExample(tt.args.astExampleString)
			want := []byte(tt.want)
			if !reflect.DeepEqual(ex.value, want) {
				t.Errorf("newBasicExample() = %v, want %v", ex.value, want)
			}
		})
	}
}
