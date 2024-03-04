package jsoac

import (
	"testing"
)

func Test_exclusiveMaximum(t *testing.T) {
	tests := []testConverterData{
		// exclusiveMaximum: true
		{
			`9 // { type: "integer", max: 10, exclusiveMaximum: true }`,
			`{
				"type": "integer",
				"example": 9,
				"maximum": 10,
				"exclusiveMaximum": true
			}`,
		},
		{
			`-11 // { type: "integer", max: -10, exclusiveMaximum: true }`,
			`{
				"type": "integer",
				"example": -11,
				"maximum": -10,
				"exclusiveMaximum": true
			}`,
		},
		{
			`9.999 // { type: "float", max: 10, exclusiveMaximum: true }`,
			`{
				"type": "number",
				"example": 9.999,
				"maximum": 10,
				"exclusiveMaximum": true
			}`,
		},
		{
			`-10.001 // { type: "float", max: -10, exclusiveMaximum: true }`,
			`{
				"type": "number",
				"example": -10.001,
				"maximum": -10,
				"exclusiveMaximum": true
			}`,
		},
		{
			`-9.99 // { type: "decimal", precision: 2, max: 10, exclusiveMaximum: true }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": -9.99,
				"maximum": 10,
				"exclusiveMaximum": true
			}`,
		},
		{
			`-10.01 // { type: "decimal", precision: 2, max: -10, exclusiveMaximum: true }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": -10.01,
				"maximum": -10,
				"exclusiveMaximum": true
			}`,
		},

		// exclusiveMaximum: false
		{
			`10 // { type: "integer", max: 10, exclusiveMaximum: false }`,
			`{
				"type": "integer",
				"example": 10,
				"maximum": 10
			}`,
		},
		{
			`-10 // { type: "integer", max: -10, exclusiveMaximum: false }`,
			`{
				"type": "integer",
				"example": -10,
				"maximum": -10
			}`,
		},
		{
			`10.000 // { type: "float", max: 10, exclusiveMaximum: false }`,
			`{
				"type": "number",
				"example": 10.000,
				"maximum": 10
			}`,
		},
		{
			`-10.000 // { type: "float", max: -10, exclusiveMaximum: false }`,
			`{
				"type": "number",
				"example": -10.000,
				"maximum": -10
			}`,
		},
		{
			`-10.00 // { type: "decimal", precision: 2, max: 10, exclusiveMaximum: false }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": -10.00,
				"maximum": 10
			}`,
		},
		{
			`-10.00 // { type: "decimal", precision: 2, max: -10, exclusiveMaximum: false }`,
			`{
				"type": "number",
				"multipleOf": 0.01,
				"example": -10.00,
				"maximum": -10
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
