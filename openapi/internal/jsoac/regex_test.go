package jsoac

import (
	"testing"
)

func Test_newOpenAPIRegex(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"Any string" // {regex: "[A-Za-z ]+"} `,
			`{
							"type": "string", 
							"example": "Any string",
							"pattern": "[A-Za-z ]+"
						}`,
		},
		{
			`"Any string" // {type: "string", regex: "[A-Za-z ]+"}`,
			`{
					"type": "string", 
					"example": "Any string",
					"pattern": "[A-Za-z ]+"
				}`,
		},
		//TODO - ERROR (code 1117): The "regex" constraint can't be used for the "email" type
		/*
			{
				`"info@mail.com" // {type: "email", regex: "[A-Za-z ]+"}`,
				`{
								"type": "string",
								"example": "info@mail.com",
								"format": "email",
								"pattern": "[A-Za-z ]+"
						}`,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
