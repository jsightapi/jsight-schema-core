package jsoac

import (
	"testing"
)

func Test_minItems(t *testing.T) {
	tests := []testConverterData{
		{
			`[ // {minItems: 1}
				"str" 
			]`,
			`{
				"type": "array",
				"minItems": 1,
				"items": {
					"type": "string", 
					"example": "str"
				}
			}`,
		},
		{
			`[ // { minItems: 2 }
				"str", 1 
			]`,
			`{
				"type": "array",
				"minItems": 2,
				"items": {
					"anyOf": [
						{"type": "string", "example": "str"},
						{"type":"integer", "example": 1}
					]
				} 
			}`,
		},
		{
			`[ // { type: "array", minItems: 3 }
				1, 
				2.3, 
				"abc"
			]`,
			`{
					"type": "array",
					"minItems": 3,
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
			`[ // {minItems: 1} 
				1, 2 
			]`,
			`{
				"type": "array",
				"minItems": 1,
				"items": {
					"anyOf": [
						{"type": "integer", "example": 1},
						{"type": "integer", "example": 2}
					]
				}	
			}`,
		},
		{
			`[ // { type: "array", nullable: true, minItems: 0 }
				]`,
			`{
					"type": "array",
					"minItems": 0,
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
