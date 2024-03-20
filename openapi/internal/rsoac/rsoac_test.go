package rsoac

import (
	"testing"
)

func Test_format(t *testing.T) {
	tests := []testConverterRegex{
		{
			`"test@test.ru" // { type: "email" }`,
			`{
				"type": "string", 
				"example": "test@test.ru",
				"format":"email"
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertRegexToOpenAPIConverter(t, data)
		})
	}
}
