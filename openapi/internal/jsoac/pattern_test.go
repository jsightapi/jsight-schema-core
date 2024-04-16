package jsoac

import (
	"testing"
)

func Test_pattern(t *testing.T) {
	tests := []testConverterData{
		{
			`"Any string" // { regex: "[A-Za-z ]+" }`,
			`{
				"type": "string", 
				"example": "Any string",
				"pattern": "[A-Za-z ]+"
			}`,
		},
		{
			`"Any string" // { type: "string", regex: "[A-Za-z ]+" }`,
			`{
				"type": "string", 
				"example": "Any string",
				"pattern": "[A-Za-z ]+"
			}`,
		},
		//TODO - ERROR (code 1117): The "regex" constraint can't be used for the "email" type
		/*
			{
				`"info@mail.com" // { type: "email", regex: "[A-Za-z ]+" }`,
				`{
								"type": "string",
								"example": "info@mail.com",
								"format": "email",
								"pattern": "[A-Za-z ]+"
						}`,
		},*/
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
