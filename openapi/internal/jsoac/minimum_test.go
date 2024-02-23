package jsoac

import (
	"testing"
)

func Test_minimum(t *testing.T) {
	tests := []testConverterData{
		{
			`1 // { min: 0 }`,
			`{
				"type": "integer", 
				"example": 1,
				"minimum": 0
			}`,
		},
		{
			`1.12 // { min: 0 }`,
			`{
				"type": "number", 
				"example": 1.12,
				"minimum": 0
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
