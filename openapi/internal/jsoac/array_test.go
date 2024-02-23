package jsoac

import (
	"testing"
)

func Test_array(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`[1, 2.3, "abc"]`,
			`{
				"type": "array",
				"items": [
					{"type": "integer", "example": 1},
					{"type": "number", "example": 2.3},
					{"type": "string", "example": "abc"}
				]
			}`,
		},
		{
			`[ // { type: "array" }
				1,
				2.3,
				"abc"
			]`,
			`{
				"type": "array",
				"items": [
					{"type": "integer", "example": 1},
					{"type": "number", "example": 2.3},
					{"type": "string", "example": "abc"}
				]
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
