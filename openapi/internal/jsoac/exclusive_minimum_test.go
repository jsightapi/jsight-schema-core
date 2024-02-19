package jsoac

import (
	"testing"
)

func Test_newOpenAPIExclusiveMinimum(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		// exclusiveMinimum: true
		{
			`1 // { type: "integer", min: 0, exclusiveMinimum: true }`,
			`{
				"type": "integer", 
				"example": 1,
				"minimum": 0, 
				"exclusiveMinimum": true
			}`,
		},
		{
			`0.001 // { type: "float", min: 0, exclusiveMinimum: true }`,
			`{
				"type": "number", 
				"example": 0.001,
				"minimum": 0,
				"exclusiveMinimum": true
			}`,
		},
		{
			`-0.999 // { type: "float", min: -1, exclusiveMinimum: true }`,
			`{
				"type": "number", 
				"example": -0.999,
				"minimum": -1,
				"exclusiveMinimum": true
			}`,
		},
		{
			`0.01 // { type: "decimal", precision: 2, min: 0, exclusiveMinimum: true }`,
			`{
				"type": "number",
				"multipleOf":0.01,
				"example": 0.01,
				"minimum": 0,
				"exclusiveMinimum": true
			}`,
		},

		// exclusiveMinimum: false
		{
			`0 // { type: "integer", min: 0, exclusiveMinimum: false }`,
			`{
				"type": "integer", 
				"example": 0,
				"minimum": 0 
			}`,
		},
		{
			`0.000 // { type: "float", min: 0, exclusiveMinimum: false }`,
			`{
				"type": "number", 
				"example": 0.000,
				"minimum": 0
			}`,
		},
		{
			`-1.000 // { type: "float", min: -1, exclusiveMinimum: false }`,
			`{
				"type": "number", 
				"example": -1.000,
				"minimum": -1
			}`,
		},
		{
			`0.00 // { type: "decimal", precision: 2, min: 0, exclusiveMinimum: false }`,
			`{
				"type": "number",
				"multipleOf":0.01,
				"example": 0.00,
				"minimum": 0
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
