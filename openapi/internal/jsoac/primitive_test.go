package jsoac

import (
	"testing"
)

func Test_primitive(t *testing.T) {
	tests := []testConverterData{
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
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
