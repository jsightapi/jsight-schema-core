package jsoac

import (
	"testing"
)

func Test_maxItems(t *testing.T) {
	tests := []testConverterData{
		{
			`[ // {maxItems: 1}
				"str" 
			]`,
			`{
				"type": "array",
				"maxItems": 1,
				"items": {
					"type": "string", 
					"example": "str"
				}
			}`,
		},
		{
			`[ // { maxItems: 2 }
				"str", 1 
			]`,
			`{
				"type": "array",
				"maxItems": 2,
				"items": {
					"anyOf": [
						{"type": "string", "example": "str"},
						{"type":"integer", "example": 1}
					]
				} 
			}`,
		},
		{
			`[ // { type: "array", maxItems: 3 }
				1, 
				2.3, 
				"abc"
			]`,
			`{
					"type": "array",
					"maxItems": 3,
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
			`[ // {maxItems: 3} 
				1, 2 
			]`,
			`{
				"type": "array",
				"maxItems": 3,
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
