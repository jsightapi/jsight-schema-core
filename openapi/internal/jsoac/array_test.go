package jsoac

import (
	"testing"
)

func Test_array(t *testing.T) {
	tests := []testConverterData{
		{ // TODO why "type":"string"?
			`[]`,
			`{
				"type": "array",
				"items": {},
				"maxItems": 0
			}`,
		},
		{
			`[ "str" ]`,
			`{
				"type": "array",
				"items": {
					"type": "string", 
					"example": "str"
				}
			}`,
		},
		{
			`[ // some text
				"str"
			]`,
			`{
				"type": "array",
				"items": {
					"type": "string", 
					"example": "str"
				},
				"description": "some text"
			}`,
		},
		{
			`[ // {minItems: 1, nullable: true} - some text
				"str"
			]`,
			`{
				"type": "array",
				"items": {
					"type": "string", 
					"example": "str"
				},
				"description": "some text",
				"minItems": 1,
				"nullable": true
			}`,
		},
		{
			`[ "str", 1 ]`,
			`{
				"type": "array",
				"items": {
					"anyOf": [
						{"type": "string", "example": "str"},
						{"type": "integer", "example": 1}
					]
				}
			}`,
		},
		{
			`[1, 2.3, "abc"]`,
			`{
					"type": "array",
					"items": {
						"anyOf": [
							{"type": "integer", "example": 1},
							{"type": "number", "example": 2.3},
							{"type": "string", "example": "abc"}
						]
					}
				}`,
		},

		{
			`[ // { type: "array" }
					1,
					2.3,
					"abc"
				]`,
			`{
					"type": "array",
					"items": {
						"anyOf": [
							{"type": "integer", "example": 1},
							{"type": "number", "example": 2.3},
							{"type": "string", "example": "abc"}
						]
					}
				}`,
		},
		{
			`[ 1, 2 ]`,
			`{
				"type": "array",
				"items": {
					"anyOf": [
						{"type": "integer", "example": 1},
						{"type": "integer", "example": 2}
					]
				}
			}`,
		},
		{
			`[ // { maxItems: 10, minItems: 2 }
				1, 2 
			]`,
			`{
				"type": "array",
				"minItems": 2,
				"maxItems": 10,
				"items": {
					"anyOf": [
						{"type": "integer", "example": 1},
						{"type": "integer", "example": 2}
					]
				}
			}`,
		},
		{
			`[ // { type: "array" }
				]`,
			`{
					"type": "array",
					"maxItems": 0,
					"items": {}
				}`,
		},
		{
			`[ // { type: "array", nullable: true }
				]`,
			`{
					"type": "array",
					"maxItems": 0,
					"items": {},
					"nullable": true
				}`,
		},
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
