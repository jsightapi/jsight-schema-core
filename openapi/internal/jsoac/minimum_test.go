package jsoac

import (
	"testing"
)

func Test_newOpenAPIMinimum(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`1 //{ min: 0 }`,
			`{
				"type": "integer", 
				"example": 1,
				"minimum": 0
			}`,
		},
		{
			`1.12 //{ min: 0 }`,
			`{
				"type": "number", 
				"example": 1.12,
				"minimum": 0
			}`,
		},
		{
			`0.12 //{ precision: 2, min: 0 }`,
			`{
				"type": "number", 
				"example": 0.12,
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
