package jsoac

import "testing"

func Test_objectKeyType(t *testing.T) {
	tests := []testComplexConverterData{
		{
			// The additionalProperties property in JSight is false by default.
			// But in OpenAPI additionalProperties appears as a workaround to implement
			// the missing functionality in OpenAPI 3.0.3
			`{
				"foo": "bar"
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
				"additionalProperties": false
			}`,
			[]testUserType{
				stringIDUserType,
			},
		},

		{
			// The additionalProperties property in JSight is false by default.
			// But in OpenAPI additionalProperties appears as a workaround to implement
			// the missing functionality in OpenAPI 3.0.3
			`{
					"foo": "bar",
					@stringId : "str"
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
								"example": "str"
							}
						]
					}
				}`,
			[]testUserType{
				stringIDUserType,
			},
		},

		{
			// The additionalProperties property in JSight is false by default.
			// But in OpenAPI additionalProperties appears as a workaround to implement
			// the missing functionality in OpenAPI 3.0.3
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
			// The additionalProperties property in JSight is false by default.
			// But in OpenAPI additionalProperties appears as a workaround to implement
			// the missing functionality in OpenAPI 3.0.3
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
			// The additionalProperties in OpenAPI is set to true by default.
			// Any properties are allowed.
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
			// The additionalProperties in OpenAPI is set to true by default.
			// Any properties are allowed.
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

		{
			`{ // { additionalProperties: "boolean"}
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
									"type": "boolean"
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
			`{ // { additionalProperties: "float"}
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
									"type": "number"
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
			`{ // { additionalProperties: "object"}
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
									"type": "object",
									"properties": {},
									"additionalProperties": false
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
			`{ // { additionalProperties: "null"}
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
									"enum": [null], 
									"example": null
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

		{
			`{ // { additionalProperties: "datetime"}
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
									"format": "date-time"
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
			`{ // { additionalProperties: "email"}
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
									"format": "email"
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
			`{ // { additionalProperties: "uri"}
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
									"format": "uri"
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
			`{ // { additionalProperties: "uuid"}
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
									"format": "uuid"
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
