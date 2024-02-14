package jsoac

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_newOpenAPIRegexRule(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
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
				`
					"Any string" // {type: "string", regex: "[A-Za-z ]+"} 
				`,
				`
					{
						"type": "string", 
						"example": "Any string",
						"pattern": "[A-Za-z ]+"
					}
				`,
			},
			//FIXME - ERROR (code 1117): The "regex" constraint can't be used for the "email" type
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
				jsightToOpenAPIRegexRule(t, tt.jsight, tt.openapi)
			})
		}
	})
}

func jsightToOpenAPIRegexRule(t *testing.T, jsight string, openapi string) {
	j := jschema.New("TestSchemaName", jsight)
	err := j.Check()
	require.NoError(t, err)

	o := New(j)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	require.JSONEq(t, openapi, string(json), "Actual: "+string(json))
}
