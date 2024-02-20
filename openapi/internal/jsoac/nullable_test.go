package jsoac

import (
	"testing"
)

func Test_newOpenAPINullable(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"foo" // { nullable: true }`,
			`{
				"type": "string",
  				"example": "foo",
				"nullable": true
			}`,
		},
		{
			`1 // { nullable: true }`,
			`{
				"type": "integer",
  				"example": 1,
				"nullable": true
			}`,
		},
		{
			`1.12 // { nullable: true }`,
			`{
				"type": "number",
  				"example": 1.12,
				"nullable": true
			}`,
		},
		{
			`1.12 // { type: "decimal", precision: 2, nullable: true }`,
			`{
				"type": "number",
  				"example": 1.12,
				"multipleOf": 0.01,
				"nullable": true
			}`,
		},
		{
			`true // { nullable: true }`,
			`{
				"type": "boolean",
  				"example": true,
				"nullable": true
			}`,
		},
		{
			`"foo" // { nullable: false }`,
			`{
				"type": "string",
  				"example": "foo"
			}`,
		},
		{
			`"test@example.com" // { type: "email", nullable: true }`,
			`{
				"type": "string",
				"format": "email",
  				"example": "test@example.com",
				"nullable": true
			}`,
		},
		{
			`"2024-02-19" // { type: "date", nullable: true }`,
			`{
				"type": "string",
				"format": "date",
  				"example": "2024-02-19",
				"nullable": true
			}`,
		},
		{
			`null // { type: "null", nullable: true }`,
			`{
				"enum": [null], 
				"example": null,
				"nullable": true
			}`,
		},
		//TODO: Enum tests
		//TODO: Array tests
		//TODO: Objects tests
		//TODO: @UserType tests
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
