package jsoac

import (
	"testing"
)

func Test_multipleOf(t *testing.T) {
	tests := []testConverterData{
		{
			`0.12 // { precision: 2 }`,
			`{
				"type": "number", 
				"example": 0.12,
				"multipleOf": 0.01
			}`,
		},
		{
			`0.9 // { precision: 1 }`,
			`{
				"type": "number", 
				"example": 0.9,
				"multipleOf": 0.1
			}`,
		},
		{
			`0.12 // { type: "decimal", precision: 2 }`,
			`{
				"type": "number", 
				"example": 0.12,
				"multipleOf": 0.01
			}`,
		},
		{
			`0.9 // { type: "decimal", precision: 1 }`,
			`{
				"type": "number", 
				"example": 0.9,
				"multipleOf": 0.1
			}`,
		},
		{
			`0.901 // { type: "decimal", precision: 3 }`,
			`{
				"type": "number", 
				"example": 0.901,
				"multipleOf": 0.001
			}`,
		},
		{
			`0.900 // { type: "decimal", precision: 1 }`,
			`{
				"type": "number", 
				"example": 0.900,
				"multipleOf": 0.1
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
