package jsoac

import (
	"testing"
)

func Test_maxLength(t *testing.T) {
	tests := []testConverterData{
		{
			`"Any string" // { maxLength: 255 }`,
			`{
				"type": "string", 
				"example": "Any string",
				"maxLength": 255
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
