package jsoac

import "testing"

func Test_type(t *testing.T) {
	tests := []testComplexConverterData{
		{
			`{} // { type: "object" }`,
			`{"type":"object","properties":{},"additionalProperties":false}`,
			[]testUserType{},
		},
		{
			`[] // { type: "array" }`,
			`{"type":"array","items":{},"maxItems":0}`,
			[]testUserType{},
		},
		{
			`123 // { type: "integer" }`,
			`{"type":"integer","example":123}`,
			[]testUserType{},
		},
		{
			`12.3 // { type: "float" }`,
			`{"type":"number","example":12.3}`,
			[]testUserType{},
		},
		{
			`12.3 // { type: "decimal", precision: 1 }`,
			`{"type":"number","example":12.3,"multipleOf":0.1}`,
			[]testUserType{},
		},
		{
			`true // { type: "boolean" }`,
			`{"type":"boolean","example":true}`,
			[]testUserType{},
		},
		{
			`"abc" // { type: "string" }`,
			`{"type":"string","example":"abc"}`,
			[]testUserType{},
		},
		{
			`"tom@cats.com" // { type: "email" }`,
			`{"type":"string","example":"tom@cats.com","format":"email"}`,
			[]testUserType{},
		},
		{
			`"https://mysite.com" // { type: "uri" }`,
			`{"type":"string","example":"https://mysite.com","format":"uri"}`,
			[]testUserType{},
		},
		{
			`"2006-01-02" // { type: "date" }`,
			`{"type":"string","example":"2006-01-02","format":"date"}`,
			[]testUserType{},
		},
		{
			`"2006-01-02T15:04:05+07:00" // { type: "datetime" }`,
			`{"type":"string","example":"2006-01-02T15:04:05+07:00","format":"date-time"}`,
			[]testUserType{},
		},
		{
			`"550e8400-e29b-41d4-a716-446655440000" // { type: "uuid" }`,
			`{"type":"string","example":"550e8400-e29b-41d4-a716-446655440000","format":"uuid"}`,
			[]testUserType{},
		},
		{
			`"white" // { type: "enum", enum: ["white", "blue", "red"] }`,
			`{"example":"white","enum":["white","blue","red"]}`,
			[]testUserType{},
		},
		{
			`"abc" // { type: "mixed", or: [{type: "integer"}, {type: "string"}] }`,
			`{"anyOf":[{"type":"integer"},{"type":"string"}],"example":"abc"}`,
			[]testUserType{},
		},
		{
			`"abc" // { type: "any" }`,
			`{"example":"abc"}`,
			[]testUserType{},
		},
		{
			`null // { type: "null" }`,
			`{"example":null,"enum":[null]}`,
			[]testUserType{},
		},
		{
			`"abc" // { type: "@stringId" }`,
			`{"allOf":[{"$ref":"#/components/schemas/stringId"}],"example":"abc"}`,
			[]testUserType{
				stringIDUserType,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIComplexConverter(t, data)
		})
	}
}
