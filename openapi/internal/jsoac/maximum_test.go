package jsoac

import (
	"testing"
)

func Test_maximum(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`1 // { max: 10 }`,
			`{
				"type": "integer", 
				"example": 1,
				"maximum": 10
			}`,
		},
		{
			`1.12 // { max: 3.4 }`,
			`{
				"type": "number", 
				"example": 1.12,
				"maximum": 3.4
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
