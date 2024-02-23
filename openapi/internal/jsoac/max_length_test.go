package jsoac

import (
	"testing"
)

func Test_maxLength(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"Any string" // { maxLength: 255 }`,
			`{
				"type": "string", 
				"example": "Any string",
				"maxLength": 255
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
