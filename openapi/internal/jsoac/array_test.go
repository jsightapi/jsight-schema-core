package jsoac

import (
	"testing"
)

func Test_array(t *testing.T) {
	tests := []testConverterData{
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
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
