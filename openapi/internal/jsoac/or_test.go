package jsoac

import (
	"testing"
)

func Test_or(t *testing.T) {
	tests := []testComplexConverterData{
		{
			`"abc" /* {or: [
					{type: "string" , maxLength: 3}, 
					{type: "integer", min: 0}
			]} */`,
			`{
				"anyOf": [
					{ "type": "string", "maxLength": 3 },
					{ "type": "integer", "minimum": 0 }
				],
				"example": "abc"
			}`,
			[]testUserType{},
		},
		{
			`"cat@mail.com" // {or: ["uuid", "email"]}`,
			`{
				"anyOf": [
					{ "type": "string", "format": "uuid" },
					{ "type": "string", "format": "email" }
				],
				"example": "cat@mail.com"
			}`,
			[]testUserType{},
		},
		{
			`"cat@mail.com" // {type: "mixed", or: ["email", "integer"]}`,
			`{
				"anyOf": [
					{ "type": "string", "format": "email" },
					{ "type": "integer" }
				],
				"example": "cat@mail.com"
			}`,
			[]testUserType{},
		},
		{
			`"cat@mail.com" // {or: ["uuid", "email"], nullable: true} - some text`,
			`{
				"anyOf": [
					{ "type": "string", "format": "uuid" },
					{ "type": "string", "format": "email" }
				],
				"example": "cat@mail.com",
				"nullable": true,
				"description": "some text"
			}`,
			[]testUserType{},
		},
		//{
		//	`123 // {or: [ {min: 100}, {type: "string"} ]}`, // TODO SERV-355
		//	`{
		//		"anyOf": [
		//			{ "type": "integer", "minimum": 100 }
		//			{ "type": "string" }
		//		],
		//		"example": 123
		//	}`,
		//	[]testUserType{},
		//},
		{
			`"abc" // {or: [ "string", "object", "array" ]}`,
			`{
				"anyOf": [
					{ "type": "string" },
					{ "type": "object", "properties": {}, "additionalProperties": false },
					{ "type": "array", "items": {}, "maxItems": 0 }
				],
				"example": "abc"
			}`,
			[]testUserType{},
		},
		{
			`1.2 // {or: [ {type: "decimal", precision: 1}, {type: "object"}, {type: "array"} ]}`,
			`{
				"anyOf": [
					{ "type": "number", "multipleOf": 0.1 },
					{ "type": "object", "properties": {}, "additionalProperties": false },
					{ "type": "array", "items": {}, "maxItems": 0 }
				],
				"example": 1.2
			}`,
			[]testUserType{},
		},

		// User types
		{
			`123 // {or: [ {type: "@stringId"}, {type: "@integerId"} ]}`,
			`{
				"anyOf": [
					{ "$ref": "#/components/schemas/stringId" },
					{ "$ref": "#/components/schemas/integerId" }
				],
				"example": 123
			}`,
			[]testUserType{
				stringIDUserType,
				integerIDUserType,
			},
		},
		{
			`123 // {or: ["@stringId", "@integerId"]}`,
			`{
				"anyOf": [
					{ "$ref": "#/components/schemas/stringId" },
					{ "$ref": "#/components/schemas/integerId" }
				],
				"example": 123
			}`,
			[]testUserType{
				stringIDUserType,
				integerIDUserType,
			},
		},
		{
			`123 // {or: [ "@stringId", {type: "integer"} ]}`,
			`{
				"anyOf": [
					{ "$ref": "#/components/schemas/stringId" },
					{ "type": "integer" }
				],
				"example": 123
			}`,
			[]testUserType{
				stringIDUserType,
				integerIDUserType,
			},
		},
		{
			`@stringId | @integerId`,
			`{
				"anyOf": [
					{ "$ref": "#/components/schemas/stringId" },
					{ "$ref": "#/components/schemas/integerId" }
				]
			}`,
			[]testUserType{
				stringIDUserType,
				integerIDUserType,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
