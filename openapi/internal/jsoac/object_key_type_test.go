package jsoac

import "testing"

func Test_objectKeyType(t *testing.T) {
	tests := []testComplexConverterData{
		{
			// The additionalProperties property in JSight is false by default. But in OpenAPI additionalProperties appears as a workaround to implement the missing functionality in OpenAPI 3.0.3
			`{
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		{
			// The additionalProperties property in JSight is false by default. But in OpenAPI additionalProperties appears as a workaround to implement the missing functionality in OpenAPI 3.0.3
			`{ // { additionalProperties: false }
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		{
			`{ // { additionalProperties: true }
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			// The additionalProperties in OpenAPI is set to true by default. Any properties are allowed.
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"]
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		{
			`{ // { additionalProperties: "any" }
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			// The additionalProperties in OpenAPI is set to true by default. Any properties are allowed.
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"]
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		{
			`{ // { additionalProperties: "string" }
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			// TODO remove duplication of the "string" type inside anyOf?
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"type": "string"
						},
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		{
			`{ // { additionalProperties: "integer"}
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			// TODO remove duplication of the "integer" type inside anyOf?
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"type": "integer"
						},
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		{
			`{ // { additionalProperties: "array"}
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"type": "array",
							"items": {}
						},
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		// TODO {additionalProperties: "boolean"}
		// TODO {additionalProperties: "float"}
		// TODO {additionalProperties: "object"}
		// TODO {additionalProperties: "null"}

		{
			`{ // { additionalProperties: "date"}
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"type": "string",
							"format": "date"
						},
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catEmailUserType,
				dogEmailUserType,
			},
		},
		// TODO {additionalProperties: "datetime"}
		// TODO {additionalProperties: "email"}
		// TODO {additionalProperties: "uri"}
		// TODO {additionalProperties: "uuid"}

		{
			`{ // { additionalProperties: "@cat"}
				"foo": "bar",
				@catEmail : 1,
				@dogEmail : "str"
			}`,
			`{
				"type": "object",
				"properties": {
					"foo": {
						"type": "string",
						"example": "bar"
					}
				},
				"required": ["foo"],
				"additionalProperties": {
					"anyOf": [
						{
							"$ref": "#/components/schemas/cat"
						},
						{
							"type": "integer",
							"example": 1
						},
						{
							"type": "string",
							"example": "str"
						}
					]
				}
			}`,
			[]testUserType{
				catUserType,
				catEmailUserType,
				dogEmailUserType,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
