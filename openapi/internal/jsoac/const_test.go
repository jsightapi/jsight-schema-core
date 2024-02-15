package jsoac

import (
	"testing"
)

func Test_newOpenAPIConst(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"OK" // {const: true}`,
			`{
							"type": "string", 
							"example":"OK",
							"enum": ["OK"]
						}`,
		},
		{
			`"OK" // {const: false}`,
			`{
							"type": "string", 
							"example":"OK"
						}`,
		},
		{
			`true // {const: true}`,
			`{
							"type": "boolean", 
							"example": true,
							"enum": [true]
						}`,
		},
		{
			`"2024-02-15" // {type: "date", const: true}`,
			`{
							"type": "string",
							"format": "date",
							"example": "2024-02-15",
							"enum": ["2024-02-15"]
						}`,
		},
		{
			`123 // {const: true}`,
			`{
							"type": "integer",
							"example": 123,
							"enum": [123]
						}`,
		},
		{
			`123.12 // {const: true}`,
			`{
							"type": "number",
							"example": 123.12,
							"enum": [123.12]
						}`,
		},
		// TODO null with const
		/*{
			`null // {const: true}`,
			`{
				"example": null,
				"enum": [null]
			}`,
		},*/
		// TODO decimal with precision const
		/*{
			`0.12 // {type: "decimal", precision: 2}`,
			`{
				"type": "number",
				"example": 0.12,
				"required": true,
				"enum": [
					0.12
				]
			}`,
		},*/
		// TODO enum const test
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}