package jsoac

import (
	"testing"
)

func Test_nullable(t *testing.T) {
	tests := []testConverterData{
		//TODO: Any tests
		//TODO: Array tests
		//TODO: Mixed tests
		//TODO: Objects tests
		//TODO: @UserType tests
		{
			`true // { nullable: true }`,
			`{
				"type": "boolean",
  				"example": true,
				"nullable": true
			}`,
		},
		{
			`"2024-02-19" // { type: "date", nullable: true }`,
			`{
				"type": "string",
				"format": "date",
  				"example": "2024-02-19",
				"nullable": true
			}`,
		},
		{
			`"2024-02-14T09:14:28+03:00" // { type: "datetime", nullable: true }`,
			`{
				"type": "string",
				"example": "2024-02-14T09:14:28+03:00",
				"format": "date-time",
				"nullable": true
			}`,
		},
		{
			`1.12 // { type: "decimal", precision: 2, nullable: true }`,
			`{
				"type": "number",
  				"example": 1.12,
				"multipleOf": 0.01,
				"nullable": true
			}`,
		},
		{
			`"test@example.com" // { type: "email", nullable: true }`,
			`{
				"type": "string",
				"format": "email",
  				"example": "test@example.com",
				"nullable": true
			}`,
		},
		{
			`"white" // { nullable: true, enum: ["white", "blue", "red"]}`,
			`{
				"example": "white",
				"enum": ["white", "blue", "red"],
				"nullable": true
			}`,
		},
		{
			`-2.1 // { nullable: true, enum: [-3, -2.1, 1.2, true, false, "-3", "0", "1.2", "string", "true"]}`,
			`{
				"example": -2.1,
				"enum": [-3, -2.1, 1.2, true, false, "-3", "0", "1.2", "string", "true"],
				"nullable": true
			}`,
		},
		{
			`null // { nullable: true, enum: ["white", "blue", "red", null]}`,
			`{
				"example": null,
				"enum": ["white", "blue", "red", null],
				"nullable": true
			}`,
		},
		{
			`null // { nullable: true, enum: ["white", "blue", "red"]}`,
			`{
				"example": null,
				"enum": ["white", "blue", "red"],
				"nullable": true
			}`,
		},
		{
			`12.34 // { type: "float", nullable: true }`,
			`{
				"type": "number", 
				"example": 12.34,
				"nullable": true
			}`,
		},
		{
			`1 // { type: "integer", nullable: true }`,
			`{
				"type": "integer",
  				"example": 1,
				"nullable": true
			}`,
		},
		{
			`null // { type: "null", nullable: true }`,
			`{
						"enum": [null],
						"example": null,
						"nullable": true
					}`,
		},
		{
			`"foo" // { type: "string", nullable: true }`,
			`{
				"type": "string",
  				"example": "foo",
				"nullable": true
			}`,
		},
		{
			`"https://www.com" // { type: "uri", nullable: true }`,
			`{
				"type": "string", 
				"example": "https://www.com",
				"format":"uri",
				"nullable": true
			}`,
		},
		{
			`"53496d7f-1374-4368-a829-74ccd47aec1c" // { type: "uuid", nullable: true }`,
			`{
				"type": "string", 
				"example": "53496d7f-1374-4368-a829-74ccd47aec1c",
				"format":"uuid",
				"nullable": true
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
