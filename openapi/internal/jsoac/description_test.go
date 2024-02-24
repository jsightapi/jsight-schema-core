package jsoac

import (
	"testing"
)

func Test_description(t *testing.T) {
	tests := []testConverterData{
		{
			`1 // {min: -99, max: 99} - Some note.`,
			`{
				"type": "integer",
				"example": 1,
				"minimum": -99,
				"maximum": 99,
				"description": "Some note."
			}`,
		},
		{
			`1 // Some note.`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Some note."
			}`,
		},
		{
			`1  /* 
            	Some note.
			*/`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Some note."
			}`,
		},
		{
			`1  /* Multiline
						 annotation
						 in several lines. 
					  */`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Multiline annotation in several lines."
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.jsight, func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}