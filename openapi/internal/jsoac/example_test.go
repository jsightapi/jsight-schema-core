package jsoac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{
			args{`"quoted string"`},
			`"\"quoted string\""`,
		},
		{
			args{`'single quoted string'`},
			`"'single quoted string'"`,
		},
		{
			args{`\*\/ \*\/`},
			`"\\*\\/ \\*\\/"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			ex := newExample(tt.args.astExampleString, true) //newStringExample(tt.args.astExampleString)
			actual, err := ex.MarshalJSON()
			require.NoError(t, err)

			expected := []byte(tt.want)

			assert.Equal(t, expected, actual, fmt.Sprintf("Expected: %s\nActual: %s\n", expected, actual))
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
		{
			args{`"quoted string"`},
			`"quoted string"`,
		},
		{
			args{`'single quoted string'`},
			`'single quoted string'`,
		},
		{
			args{`\*\/ \*\/`},
			`\*\/ \*\/`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			ex := newExample(tt.args.astExampleString, false)
			actual, err := ex.MarshalJSON()
			require.NoError(t, err)

			expected := []byte(tt.want)

			assert.Equal(t, expected, actual, fmt.Sprintf("Expected: %s\nActual: %s\n", expected, actual))
		})
	}
}
