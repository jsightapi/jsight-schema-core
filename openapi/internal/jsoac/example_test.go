package jsoac

import (
	"reflect"
	"testing"
)

func Test_example_string(t *testing.T) {
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
			ex := newExample(tt.args.astExampleString, true) //newStringExample(tt.args.astExampleString)
			want := []byte(tt.want)
			if !reflect.DeepEqual(ex.value, want) {
				t.Errorf("newStringExample() = %s, want %s", ex.value, want)
			}
		})
	}
}

func Test_example(t *testing.T) {
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
			ex := newExample(tt.args.astExampleString, false)
			want := []byte(tt.want)
			if !reflect.DeepEqual(ex.value, want) {
				t.Errorf("newExample() = %s, want %s", ex.value, want)
			}
		})
	}
}
