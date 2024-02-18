package jsoac

import (
	"testing"
)

func Test_Object(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`{
				"str": "some string",
				"int": 123,
				"num": 12.3,
				"bool": true,
				"arr": [1,2],
				"obj": {"key": "val"}
			}`,
			`{
				"type": "object",
				"properties": {
					"str": {"type": "string", "example": "some string"},
					"int": {"type": "integer", "example": 123},
					"num": {"type": "number", "example": 12.3},
					"bool": {"type": "boolean", "example": true},
					"arr": {
						"type": "array",
						"items": [
							{"type": "integer", "example": 1},
							{"type": "integer", "example": 2}
						]
					},
					"obj": {
						"type": "object",
						"properties": {
							"key": {"type": "string", "example": "val"}
						}
					}
				}
			}`,
		},
		{
			`{ // { type: "object" }
				"str": "abc"
			}`,
			`{
				"type": "object",
				"properties": {
					"str": {"type": "string", "example": "abc"}
				}
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
