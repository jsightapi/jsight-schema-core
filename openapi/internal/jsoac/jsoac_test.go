package jsoac

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

func TestNewOpenAPI(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		tests := []struct {
			jsight  string
			openapi string
		}{
			{
				`"some string"`,
				`{"type": "string", "example": "some string"}`,
			},
			{
				`123`,
				`{"type": "integer", "example": 123}`,
			},
			{
				`123.4`,
				`{"type": "number", "example": 123.4}`,
			},
			{
				`true`,
				`{"type": "boolean", "example": true}`,
			},
			{
				`false`,
				`{"type": "boolean", "example": false}`,
			},
			{
				`null`,
				`{"enum": [null], "example": null}`,
			},
			{
				`[1, 2.3, "abc"]`,
				`{
				"type": "array",
				"items": [
					{"type": "integer", "example": 1},
					{"type": "number", "example": 2.3},
					{"type": "string", "example": "abc"}
				]
			}`,
			},
			{
				`{
				"str": "some string",
				"int": 123,
				"num": 12.3,
				"bool": true,
				"arr": [1,2],
				"obj": {"key": "val"}
				}`,
				`{
				"type": "object",
				"properties": {
					"str": {"type": "string", "example": "some string"},
					"int": {"type": "integer", "example": 123},
					"num": {"type": "number", "example": 12.3},
					"bool": {"type": "boolean", "example": true},
					"arr": {
						"type": "array",
						"items": [
							{"type": "integer", "example": 1},
							{"type": "integer", "example": 2}
						]
					},
					"obj": {
						"type": "object",
						"properties": {
							"key": {"type": "string", "example": "val"}
						}
					}
				}
				}`,
			},

			{
				`
					"test@test.ru" //{ type: "email" }
				`,
				`
					{
						"type": "string", 
						"example": "test@test.ru",
						"format":"email"
					}
				`,
			},
			{
				`
					"https://www.com" //{ type: "uri" }
				`,
				`
					{
						"type": "string", 
						"example": "https://www.com",
						"format":"uri"
					}
				`,
			},
			{
				`
					"53496d7f-1374-4368-a829-74ccd47aec1c" //{ type: "uuid" }
				`,
				`
					{
						"type": "string", 
						"example": "53496d7f-1374-4368-a829-74ccd47aec1c",
						"format":"uuid"
					}
				`,
			},
			{
				`
					"2024-02-14" //{ type: "date" }
				`,
				`
					{
						"type": "string", 
						"example": "2024-02-14",
						"format":"date"
					}
				`,
			},
			{
				`
					"2024-02-14T09:14:28+03:00" //{ type: "datetime" }
				`,
				`
					{
						"type": "string", 
						"example": "2024-02-14T09:14:28+03:00",
						"format":"date-time"
					}
				`,
			},
			//TODO: support for different integer formats, such as int32 and int64
			/*{
				`123 // {type: "integer"}`,
				`{"type": "integer", "example": 123, "format": "int32"}`,
			},*/
			{
				`123 // {type: "integer"}`,
				`{"type": "integer", "example": 123, "format": "int32"}`,
			},
		}

		for _, tt := range tests {
			t.Run(tt.jsight, func(t *testing.T) {
				jsightToOpenAPI(t, tt.jsight, tt.openapi)
			})
		}
	})
}

func jsightToOpenAPI(t *testing.T, jsight string, openapi string) {
	j := jschema.New("TestSchemaName", jsight)
	err := j.Check()
	require.NoError(t, err)

	o := New(j)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	require.JSONEq(t, openapi, string(json), "Actual: "+string(json))
}
