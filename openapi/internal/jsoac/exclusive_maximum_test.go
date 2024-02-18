package jsoac

import (
	"testing"
)

func Test_newOpenAPIExclusiveMaximum(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`1 // { max: 10, exclusiveMaximum: true }`,
			`{
				"type": "integer", 
				"example": 1,
				"maximum": 10, 
				"exclusiveMaximum": true
			}`,
		},
		{
			`1.12 // { max: 10, exclusiveMaximum: true }`,
			`{
				"type": "number", 
				"example": 1.12,
				"maximum": 10,
				"exclusiveMaximum": true
			}`,
		},
		{
			`0.12 // { max: 10, exclusiveMaximum: true }`,
			`{
				"type": "number", 
				"example": 0.12,
				"maximum": 10,
				"exclusiveMaximum": true
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
