package jsoac

import (
	"testing"
)

func Test_allOf(t *testing.T) {
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
				catUserType,
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
				catUserType,
				dogUserType,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
