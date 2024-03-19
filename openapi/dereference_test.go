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

var refToObj = testUserType{
	"@refToObj",
	`@obj`,
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
			[]testUserType{refToObj, obj},
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
			`123 // {type: "any"}`,
			[]testUserType{objOrArr, obj, arr},
			[]SchemaInfoType{SchemaInfoTypeAny},
		},
		{
			`123 // { or: [ {type: "integer"}, {type: "string"} ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		// TODO SERV-355
		// ERROR (code 906): Type is required inside the "or" rule
		//{
		//	`123 // {or: [ {minLength: 1}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		//},
		//{
		//	`123 // {or: [ {minItems: 0}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]SchemaInfoType{SchemaInfoTypeArray, SchemaInfoTypeScalar},
		//},
		//{
		//	`123 // {or: [ {additionalProperties: true}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeScalar},
		//},
		{
			`123 // { or: [ "integer", "string" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ "uuid", "integer" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ {type: "email"}, "integer"] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ {type: "null"}, "integer" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ "any", "integer" ] }`,
			[]testUserType{},
			[]SchemaInfoType{SchemaInfoTypeAny, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ "@obj", "@arr", "@refToObj", "integer" ] }`,
			[]testUserType{obj, arr, refToObj},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeArray, SchemaInfoTypeObject, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ {type: "@obj"}, {type: "@arr"}, {type: "@refToObj"}, "integer" ] }`,
			[]testUserType{obj, arr, refToObj},
			[]SchemaInfoType{SchemaInfoTypeObject, SchemaInfoTypeArray, SchemaInfoTypeObject, SchemaInfoTypeScalar},
		},
		{
			`123 // { or: [ {type: "integer", min: 0}, "string", "@obj", {type: "@arr"}, "@refToObj" ] }`,
			[]testUserType{obj, arr, refToObj},
			[]SchemaInfoType{SchemaInfoTypeScalar, SchemaInfoTypeScalar, SchemaInfoTypeObject, SchemaInfoTypeArray, SchemaInfoTypeObject},
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
