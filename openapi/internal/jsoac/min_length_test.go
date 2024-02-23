package jsoac

import (
	"testing"
)

func Test_minLength(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"Any string" // { minLength: 3 }`,
			`{
				"type": "string", 
				"example": "Any string",
				"minLength": 3
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
