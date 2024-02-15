package jsoac

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_newOpenAPIRules(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		tests := []struct {
			jsight  string
			openapi string
		}{
			/* REGEX TESTS */
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

			/* TYPE's TESTS */
			{
				`"test@test.ru" //{ type: "email" }`,
				`{
							"type": "string", 
							"example": "test@test.ru",
							"format":"email"
						}`,
			},
			{
				`"https://www.com" //{ type: "uri" }`,
				`{
							"type": "string", 
							"example": "https://www.com",
							"format":"uri"
						}`,
			},
			{
				`"53496d7f-1374-4368-a829-74ccd47aec1c" //{ type: "uuid" }`,
				`{
							"type": "string", 
							"example": "53496d7f-1374-4368-a829-74ccd47aec1c",
							"format":"uuid"
						}`,
			},
			{
				`"2024-02-14" //{ type: "date" }`,
				`{
							"type": "string", 
							"example": "2024-02-14",
							"format":"date"
						}`,
			},
			{
				`"2024-02-14T09:14:28+03:00" //{ type: "datetime" }`,
				`{
							"type": "string", 
							"example": "2024-02-14T09:14:28+03:00",
							"format":"date-time"
						}`,
			},
			{
				`123 // {type: "integer"}`,
				`{"type": "integer", "example": 123}`,
			},

			/* CONST TESTS */
			{
				`"OK" // {const: true}`,
				`{
							"type": "string", 
							"example":"OK",
							"required": true,
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
							"required": true,
							"enum": [true]
						}`,
			},
			{
				`"2024-02-15" // {type: "date", const: true}`,
				`{
							"type": "string",
							"format": "date",
							"example": "2024-02-15",
							"required": true,
							"enum": ["2024-02-15"]
						}`,
			},
			{
				`123 // {const: true}`,
				`{
							"type": "integer",
							"example": 123,
							"required": true,
							"enum": [123]
						}`,
			},
			{
				`123.12 // {const: true}`,
				`{
							"type": "number",
							"example": 123.12,
							"required": true,
							"enum": [123.12]
						}`,
			},
			// TODO null with const
			/*{
				`null // {const: true}`,
				`{
					"example": null,
					"required": true,
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
				jsightToOpenAPIRules(t, tt.jsight, tt.openapi)
			})
		}
	})
}

func jsightToOpenAPIRules(t *testing.T, jsight string, openapi string) {
	j := jschema.New("TestSchemaName", jsight)
	err := j.Check()
	require.NoError(t, err)

	o := New(j)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	require.JSONEq(t, openapi, string(json), "Actual: "+string(json))
}
