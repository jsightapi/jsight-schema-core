package jsoac

import (
	"testing"
)

func Test_newOpenAPIFormat(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"test@test.ru" // { type: "email" }`,
			`{
				"type": "string", 
				"example": "test@test.ru",
				"format":"email"
			}`,
		},
		{
			`"https://www.com" // { type: "uri" }`,
			`{
				"type": "string", 
				"example": "https://www.com",
				"format":"uri"
			}`,
		},
		{
			`"53496d7f-1374-4368-a829-74ccd47aec1c" // { type: "uuid" }`,
			`{
				"type": "string", 
				"example": "53496d7f-1374-4368-a829-74ccd47aec1c",
				"format":"uuid"
			}`,
		},
		{
			`"2024-02-14" // { type: "date" }`,
			`{
				"type": "string", 
				"example": "2024-02-14",
				"format":"date"
			}`,
		},
		{
			`"2024-02-14T09:14:28+03:00" // { type: "datetime" }`,
			`{
				"type": "string", 
				"example": "2024-02-14T09:14:28+03:00",
				"format":"date-time"
			}`,
		},
		{
			`12.34 // { type: "float" }`,
			`{"type": "number", "format":"float", "example": 12.34}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
