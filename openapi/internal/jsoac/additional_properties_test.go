package jsoac

import (
	"testing"
)

func Test_additionalProperties(t *testing.T) {
	tests := []testComplexConverterData{
		// boolean
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

		// basic types
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
					"type": "string",
					"format": "email"
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
					"type": "string",
					"format": "email"
				}
			}`,
			[]testUserType{},
		},
		{
			`{ // {additionalProperties: "decimal"}
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
					"type": "string",
					"format": "email"
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
					"type": "string",
					"format": "email"
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
					"type": "string",
					"format": "email"
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
					"type": "string",
					"format": "email"
				}
			}`,
			[]testUserType{},
		},

		// formatted types
		// TODO date
		// TODO datetime
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
		// TODO date
		// TODO datetime
		// TODO uri
		// TODO uuid

		// other
		{
			`{ // {additionalProperties: "any"}
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
		// TODO enum
		// TODO mixed
		// TODO null

		// user types
		// TODO additionalProperties: "@cat"
		// TODO additionalProperties: {type: "@cat"}
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
