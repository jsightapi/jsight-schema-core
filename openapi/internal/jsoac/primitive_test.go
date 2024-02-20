package jsoac

import (
	"testing"
)

func Test_Primitive(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"some string"`,
			`{"type": "string", "example": "some string"}`,
		},
		{
			`"some string" // { type: "string" }`,
			`{"type": "string", "example": "some string"}`,
		},
		{
			`123`,
			`{"type": "integer", "example": 123}`,
		},
		{
			`123 // { type: "integer" }`,
			`{"type": "integer", "example": 123}`,
		},
		{
			`12.34`,
			`{"type": "number", "example": 12.34}`,
		},
		{
			`12.34 // { type: "float" }`,
			`{"type": "number", "example": 12.34}`,
		},
		{
			`true`,
			`{"type": "boolean", "example": true}`,
		},
		{
			`true // { type: "boolean" }`,
			`{"type": "boolean", "example": true}`,
		},
		{
			`false`,
			`{"type": "boolean", "example": false}`,
		},
		{
			`false // { type: "boolean" }`,
			`{"type": "boolean", "example": false}`,
		},
		{
			`null`,
			`{"enum": [null], "example": null}`,
		},
		{
			`null // { type: "null" }`,
			`{"enum": [null], "example": null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
