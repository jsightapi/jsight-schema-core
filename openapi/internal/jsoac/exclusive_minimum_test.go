package jsoac

import (
	"testing"
)

func Test_exclusiveMinimum(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		// exclusiveMinimum: true
		{
			`11 // { type: "integer", min: 10, exclusiveMinimum: true }`,
			`{
				"type": "integer",
				"example": 11,
				"minimum": 10,
				"exclusiveMinimum": true
			}`,
		},
		{
			`-9 // { type: "integer", min: -10, exclusiveMinimum: true }`,
			`{
				"type": "integer",
				"example": -9,
				"minimum": -10,
				"exclusiveMinimum": true
			}`,
		},
		{
			`10.001 // { type: "float", min: 10, exclusiveMinimum: true }`,
			`{
				"type": "number",
				"example": 10.001,
				"minimum": 10,
				"exclusiveMinimum": true
			}`,
		},
		{
			`-9.999 // { type: "float", min: -10, exclusiveMinimum: true }`,
			`{
				"type": "number",
				"example": -9.999,
				"minimum": -10,
				"exclusiveMinimum": true
			}`,
		},
		{
			`10.01 // { type: "decimal", precision: 2, min: 10, exclusiveMinimum: true }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": 10.01,
				"minimum": 10,
				"exclusiveMinimum": true
			}`,
		},
		{
			`-9.99 // { type: "decimal", precision: 2, min: -10, exclusiveMinimum: true }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": -9.99,
				"minimum": -10,
				"exclusiveMinimum": true
			}`,
		},

		// exclusiveMinimum: false
		{
			`10 // { type: "integer", min: 10, exclusiveMinimum: false }`,
			`{
				"type": "integer",
				"example": 10,
				"minimum": 10
			}`,
		},
		{
			`-10 // { type: "integer", min: -10, exclusiveMinimum: false }`,
			`{
				"type": "integer",
				"example": -10,
				"minimum": -10
			}`,
		},
		{
			`10.000 // { type: "float", min: 10, exclusiveMinimum: false }`,
			`{
				"type": "number",
				"example": 10.000,
				"minimum": 10
			}`,
		},
		{
			`-10.000 // { type: "float", min: -10, exclusiveMinimum: false }`,
			`{
				"type": "number",
				"example": -10.000,
				"minimum": -10
			}`,
		},
		{
			`10.00 // { type: "decimal", precision: 2, min: 10, exclusiveMinimum: false }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": 10.00,
				"minimum": 10
			}`,
		},
		{
			`-10.00 // { type: "decimal", precision: 2, min: -10, exclusiveMinimum: false }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": -10.00,
				"minimum": -10
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
