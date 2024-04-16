package jsoac

import "testing"

func Test_ref(t *testing.T) {
	tests := []testComplexConverterData{
		{
			`"1111 222222" // { type: "@passportNumber", nullable: true } - Some text`,
			`{
				"example": "1111 222222",
				"description": "Some text",
				"nullable": true,
				"allOf": [
					{ "$ref": "#/components/schemas/passportNumber" }
				]
			}`,
			[]testUserType{
				{name: "@passportNumber", jsight: `"1234 567890" // { regex: "^\\d{4} \\d{6}$" }`},
			},
		},
		{
			`"1111 222222" // { type: "@passportNumber", nullable: false }`,
			`{
				"example": "1111 222222",
				"allOf": [
					{ "$ref": "#/components/schemas/passportNumber" }
				]
			}`,
			[]testUserType{
				{name: "@passportNumber", jsight: `"1234 567890" // { regex: "^\\d{4} \\d{6}$" }`},
			},
		},
		{
			`@cat`,
			`{
				"$ref": "#/components/schemas/cat"
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`@cat // { nullable: false }`,
			`{
				"$ref": "#/components/schemas/cat"
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`@cat // { nullable: true }`,
			`{
				"nullable": true,
				"allOf": [
					{ "$ref": "#/components/schemas/cat" }
				]
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`[ @cat ]`,
			`{
				"type": "array",
				"items": {
					"$ref": "#/components/schemas/cat"
				}
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`[
				@cat // { nullable: true } - Some text
			]`,
			`{
				"type": "array",
				"items": {
					"description": "Some text",
					"nullable": true,
					"allOf": [
						{ "$ref": "#/components/schemas/cat" }
					]
				}
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`{
				"key": @cat
			}`,
			`{
				"type": "object",
				"properties": {
					"key": {
						"$ref": "#/components/schemas/cat"
					}
				},
				"additionalProperties": false,
				"required": [ "key" ]
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`{
				"key": @cat // { optional: true }
			}`,
			`{
				"type": "object",
				"properties": {
					"key": {
						"$ref": "#/components/schemas/cat"
					}
				},
				"additionalProperties": false
			}`,
			[]testUserType{
				catUserType,
			},
		},
		{
			`{
				"key": @cat // { optional: true, nullable: true }
			}`,
			`{
				"type": "object",
				"properties": {
					"key": {
						"nullable": true,
						"allOf": [
							{ "$ref": "#/components/schemas/cat" }
						]
					}
				},
				"additionalProperties": false
			}`,
			[]testUserType{
				catUserType,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
