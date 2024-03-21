package rsoac

import (
	"testing"
)

func Test_Regex(t *testing.T) {
	tests := []testConverterRegex{
		{
			`/OK/`,
			`{
        		"type": "string",
        		"pattern": "OK"
			}`,
		},
		{
			`/^[A-Z][a-z]*( [A-Z][a-z]*)*$/`,
			`{
        		"type": "string",
        		"pattern": "^[A-Z][a-z]*( [A-Z][a-z]*)*$"
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
