package openapi

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

type testDereferenceData struct {
	jsight    string
	userTypes []testUserType
	expected  []SchemaInfoType
}

func (t testDereferenceData) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

var obj = testUserType{
	"@obj",
	`{ "key": "value" }`,
}

var obj2 = testUserType{
	"@obj2",
	`{ "key": "value" }`,
}

var refToObj = testUserType{
	"@refToObj",
	`@obj2`,
}

var arr = testUserType{
	"@arr",
	`[1,2,3]`,
}

var objOrArr = testUserType{
	"@objOrArr",
	`@obj | @arr`,
}

func Test_Dereference(t *testing.T) {
	tests := []testDereferenceData{
		{
			`{ "catName": "Tom" }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeObject},
		},
		{
			`@obj`,
			[]testUserType{obj},
			[]SchemaInfoType{SchemaInfoTypeObject},
		},
		{
			`@refToObj`,
			[]testUserType{refToObj, obj2},
			[]SchemaInfoType{SchemaInfoTypeObject},
		},
		{
			`@arr`,
			[]testUserType{arr},
			[]SchemaInfoType{SchemaInfoTypeArray},
		},
		{
			`@objOrArr`,
			[]testUserType{objOrArr, obj, arr},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeArray},
		},
		{
			`null`,
			[]testUserType{objOrArr, obj, arr},
			[]SchemaInfoType{SchemaInfoTypeScalar},
		},
		{
			`1 // {type: "any"}`,
			[]testUserType{objOrArr, obj, arr},
			[]SchemaInfoType{SchemaInfoTypeAny},
		},
		{
			`2 // { or: [ {type: "integer"}, {type: "string"} ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		// TODO SERV-355
		// ERROR (code 906): Type is required inside the "or" rule
		//{
		//	`3 // {or: [ {minLength: 1}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		//},
		//{
		//	`4 // {or: [ {minItems: 0}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]SchemaInfoType{SchemaInfoTypeArray, SchemaInfoTypeScalar},
		//},
		//{
		//	`5 // {or: [ {additionalProperties: true}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeScalar},
		//},
		{
			`6 // { or: [ "integer", "string" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`7 // { or: [ "uuid", "integer" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`8 // { or: [ {type: "email"}, "integer"] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`9 // { or: [ {type: "null"}, "integer" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`10 // { or: [ "any", "integer" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeAny, SchemaInfoTypeScalar},
		},
		{
			`11 // { or: [ "@obj", "@arr", "@refToObj", "integer" ] }`,
			[]testUserType{obj, obj2, arr, refToObj},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeArray, SchemaInfoTypeObject, SchemaInfoTypeScalar},
		},
		{
			`12 // { or: [ {type: "@obj"}, {type: "@arr"}, {type: "@refToObj"}, "integer" ] }`,
			[]testUserType{obj, obj2, arr, refToObj},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeArray, SchemaInfoTypeObject, SchemaInfoTypeScalar},
		},
		{
			`13 // { or: [ {type: "integer", min: 0}, "string", "@obj", {type: "@arr"}, "@refToObj" ] }`,
			[]testUserType{obj, obj2, arr, refToObj},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar, SchemaInfoTypeObject, SchemaInfoTypeArray, SchemaInfoTypeObject},
		},
		{
			`@query1`,
			[]testUserType{
				{
					"@query1",
					`@query1 | @query2`,
				},
				{
					"@query2",
					`@query2 | @query3`,
				},
				{
					"@query3",
					`{}`,
				},
			},
			[]SchemaInfoType{SchemaInfoTypeObject},
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertDereference(t, data)
		})
	}
}

func assertDereference(t *testing.T, data testDereferenceData) {
	jSchema := buildJSchema(t, data.jsight, data.userTypes)

	informers := Dereference(jSchema)

	actual := make([]SchemaInfoType, 0, len(informers))
	for _, e := range informers {
		actual = append(actual, e.Type())
	}

	require.Equal(t, data.expected, actual)
}
