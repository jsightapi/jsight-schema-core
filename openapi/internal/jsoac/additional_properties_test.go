package jsoac

import (
	"testing"
)

func Test_additionalProperties(t *testing.T) {
	tests := []testComplexConverterData{
		// boolean value
		{
			`{
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": false
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: false}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": false
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: true}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"]
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "any"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"]
			}`,
			[]testUserType{},
		},

		// primitive types
		{
			`{ // {additionalProperties: "array"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "array",
					"items": {}
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "boolean"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "boolean"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "float"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "number"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "integer"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "integer"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "object"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "object"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "string"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "string"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "null"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"enum": [null]
				}
			}`,
			[]testUserType{},
		},

		// formatted types
		{
			`{ // {additionalProperties: "date"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "string",
					"format": "date"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "datetime"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "string",
					"format": "date-time"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "email"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "string",
					"format": "email"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "uri"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "string",
					"format": "uri"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "uuid"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"type": "string",
					"format": "uuid"
				}
			}`,
			[]testUserType{},
		},

		// user types
		{
			`{ // {additionalProperties: "@cat"}
				"foo": "bar"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": { "type": "string", "example": "bar" }
				},
				"required": ["foo"],
				"additionalProperties": {
					"$ref": "#/components/schemas/cat"
				}
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
