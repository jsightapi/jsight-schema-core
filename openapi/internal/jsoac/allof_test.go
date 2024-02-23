package jsoac

import (
	"testing"
)

func Test_newAllOf(t *testing.T) {
	tests := []testComplexConverterData{
		{
			`{ // { allOf: "@cat" }
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": false,
				"allOf": [{
					"$ref": "#/components/schemas/cat"
				}]
			}`,
			[]testUserType{
				testCatUserType,
			},
		},
		{
			`{ // { allOf: [ "@cat", "@dog" ] }
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": false,
				"allOf": [
					{
						"$ref": "#/components/schemas/cat"
					},
					{
						"$ref": "#/components/schemas/dog"
					}
				]
			}`,
			[]testUserType{
				testCatUserType,
				testDogUserType,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
