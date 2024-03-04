package jsoac

import (
	"testing"
)

func Test_primitive(t *testing.T) {
	tests := []testConverterData{
		{
			`""`,
			`{"type": "string", "example": ""}`, // TODO discuss
		},
		{
			`"some string"`,
			`{"type": "string", "example": "some string"}`,
		},
		{
			`"some string" // some text`,
			`{"type": "string", "example": "some string", "description": "some text"}`,
		},
		{
			`"some string" // {minLength: 3} - some text`,
			`{"type": "string", "example": "some string", "description": "some text", "minLength": 3}`,
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
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
