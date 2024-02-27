package jsoac

import (
	"testing"
)

func Test_object(t *testing.T) {
	tests := []testConverterData{
		{
			`{}`,
			`{
				"type": "object",
				"properties": {},
				"additionalProperties": false
			}`,
		},
		{
			`{} // some text`,
			`{
				"type": "object",
				"properties": {},
				"additionalProperties": false,
				"description": "some text"
			}`,
		},
		{
			`{} // {nullable: true} - some text`,
			`{
				"type": "object",
				"properties": {},
				"additionalProperties": false,
				"nullable": true,
				"description": "some text"
			}`,
		},
		{
			`{
				"str": "abc"
			}`,
			`{
				"type": "object",
				"properties": {
					"str": {"type": "string", "example": "abc"}
				},
				"additionalProperties": false,
				"required": [ "str" ]
			}`,
		},
		{
			`{
				"required1": 111,
				"required2": 222, // {optional: false}
				"optional1": 333  // {optional: true}
			}`,
			`{
				"type": "object",
				"properties": {
					"required1": {"type": "integer", "example": 111},
					"required2": {"type": "integer", "example": 222},
					"optional1": {"type": "integer", "example": 333}
				},
				"additionalProperties": false,
				"required": [ "required1", "required2" ]
			}`,
		},
		{
			`{ // { type: "object" }
				"str": "abc"
			}`,
			`{
				"type": "object",
				"properties": {
					"str": {"type": "string", "example": "abc"}
				},
				"additionalProperties": false,
				"required": [ "str" ]
			}`,
		},
		{
			`{
				"str": "some string",
				"int": 123,
				"num": 12.3,
				"bool": true,
				"arr": [1,2],
				"obj": {"key": "val"}
			}`,
			`{
				"type": "object",
				"properties": {
					"str": {"type": "string", "example": "some string"},
					"int": {"type": "integer", "example": 123},
					"num": {"type": "number", "example": 12.3},
					"bool": {"type": "boolean", "example": true},
					"arr": {
						"type": "array",
						"items": {
							"anyOf": [
								{"type": "integer", "example": 1},
								{"type": "integer", "example": 2}
							]
						}
					},
					"obj": {
						"type": "object",
						"properties": {
							"key": {"type": "string", "example": "val"}
						},
						"additionalProperties": false,
						"required": [ "key" ]
					}
				},
				"additionalProperties": false,
				"required": [ "str", "int", "num", "bool", "arr", "obj" ]
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
