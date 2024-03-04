package jsoac

import (
	"testing"
)

func Test_minLength(t *testing.T) {
	tests := []testConverterData{
		{
			`"Any string" // { minLength: 3 }`,
			`{
				"type": "string", 
				"example": "Any string",
				"minLength": 3
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
