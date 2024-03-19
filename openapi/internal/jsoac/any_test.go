package jsoac

import (
	"testing"
)

func Test_any(t *testing.T) {
	tests := []testConverterData{
		{
			`"some string" // {type: "any"}`,
			`{"example": "some string"}`,
		},
		{
			`"some string" // {type: "any"} - some text`,
			`{"example": "some string", "description": "some text"}`,
		},
		{
			`123 // {type: "any"}`,
			`{"example": 123}`,
		},
		{
			`12.34  // {type: "any"}`,
			`{"example": 12.34}`,
		},
		{
			`true // {type: "any"}`,
			`{"example": true}`,
		},
		{
			`false // {type: "any"}`,
			`{"example": false}`,
		},
		{
			`false // {type: "any", nullable: true}`,
			`{
				"example": false,
				"nullable": true
			}`,
		},
		{
			`null // { type: "any" }`,
			`{"example": null}`,
		},
		{
			`null // { type: "any" } - some text`,
			`{"example": null, "description": "some text"}`,
		},
		{
			`{} // { type: "any" }`,
			`{}`,
		},
		{
			`"" // { type: "any" }`,
			`{}`,
		},
		{
			`{} // { type: "any" } - some text`,
			`{"description": "some text"}`,
		},
		{
			`[] // { type: "any" }`,
			`{}`,
		},
		{
			`[] // { type: "any" } - some text`,
			`{"description": "some text"}`,
		},
		{
			`[
				1, // {type: "any"}
				"abc@def.com", // {type: "any"}
				3.14, // {type: "any"}
				true // {type: "any"}
			]`,
			`{
				"type": "array",
				"items": {
					"anyOf": [
						{"example": 1},
						{"example": "abc@def.com"},
						{"example": 3.14},
						{"example": true}
					]
				}
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
