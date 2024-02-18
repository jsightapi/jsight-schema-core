package jsoac

import (
	"testing"
)

func Test_newOpenAPIExclusiveMinimum(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`1 // { min: 0, exclusiveMinimum: true }`,
			`{
				"type": "integer", 
				"example": 1,
				"minimum": 0, 
				"exclusiveMinimum": true
			}`,
		},
		{
			`1.12 // { min: 0, exclusiveMinimum: true }`,
			`{
				"type": "number", 
				"example": 1.12,
				"minimum": 0,
				"exclusiveMinimum": true
			}`,
		},
		{
			`0.12 // { min: 0, exclusiveMinimum: true }`,
			`{
				"type": "number", 
				"example": 0.12,
				"minimum": 0,
				"exclusiveMinimum": true
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
