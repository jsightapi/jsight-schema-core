package jsoac

import (
	"testing"
)

func Test_maximum(t *testing.T) {
	tests := []testConverterData{
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
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
