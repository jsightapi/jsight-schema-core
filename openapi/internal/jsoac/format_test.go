package jsoac

import (
	"testing"
)

func Test_format(t *testing.T) {
	tests := []testConverterData{
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
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
