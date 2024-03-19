package openapi

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

type testDereferenceData struct {
	jsight    string
	userTypes []testUserType
	expected  []ElementType
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
			[]ElementType{ElementTypeObject},
		},
		{
			`@obj`,
			[]testUserType{obj},
			[]ElementType{ElementTypeObject},
		},
		{
			`@refToObj`,
			[]testUserType{refToObj, obj},
			[]ElementType{ElementTypeObject},
		},
		{
			`@arr`,
			[]testUserType{arr},
			[]ElementType{ElementTypeArray},
		},
		{
			`@objOrArr`,
			[]testUserType{objOrArr, obj, arr},
			[]ElementType{ElementTypeObject, ElementTypeArray},
		},
		{
			`null`,
			[]testUserType{objOrArr, obj, arr},
			[]ElementType{ElementTypeScalar},
		},
		{
			`123 // {type: "any"}`,
			[]testUserType{objOrArr, obj, arr},
			[]ElementType{ElementTypeAny},
		},
		{
			`123 // { or: [ {type: "integer"}, {type: "string"} ] }`,
			[]testUserType{},
			[]ElementType{ElementTypeScalar, ElementTypeScalar},
		},
		// TODO SERV-355
		// ERROR (code 906): Type is required inside the "or" rule
		//{
		//	`123 // {or: [ {minLength: 1}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]ElementType{ElementTypeScalar, ElementTypeScalar},
		//},
		//{
		//	`123 // {or: [ {minItems: 0}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]ElementType{ElementTypeArray, ElementTypeScalar},
		//},
		//{
		//	`123 // {or: [ {additionalProperties: true}, {type: "integer"} ]}`,
		//	[]testUserType{},
		//	[]ElementType{ElementTypeObject, ElementTypeScalar},
		//},
		{
			`123 // { or: [ "integer", "string" ] }`,
			[]testUserType{},
			[]ElementType{ElementTypeScalar, ElementTypeScalar},
		},
		{
			`123 // { or: [ "uuid", "integer" ] }`,
			[]testUserType{},
			[]ElementType{ElementTypeScalar, ElementTypeScalar},
		},
		{
			`123 // { or: [ {type: "email"}, "integer"] }`,
			[]testUserType{},
			[]ElementType{ElementTypeScalar, ElementTypeScalar},
		},
		{
			`123 // { or: [ {type: "null"}, "integer" ] }`,
			[]testUserType{},
			[]ElementType{ElementTypeScalar, ElementTypeScalar},
		},
		{
			`123 // { or: [ "any", "integer" ] }`,
			[]testUserType{},
			[]ElementType{ElementTypeAny, ElementTypeScalar},
		},
		{
			`123 // { or: [ "@obj", "@arr", "@refToObj", "integer" ] }`,
			[]testUserType{obj, arr, refToObj},
			[]ElementType{ElementTypeObject, ElementTypeArray, ElementTypeObject, ElementTypeScalar},
		},
		{
			`123 // { or: [ {type: "@obj"}, {type: "@arr"}, {type: "@refToObj"}, "integer" ] }`,
			[]testUserType{obj, arr, refToObj},
			[]ElementType{ElementTypeObject, ElementTypeArray, ElementTypeObject, ElementTypeScalar},
		},
		{
			`123 // { or: [ {type: "integer", min: 0}, "string", "@obj", {type: "@arr"}, "@refToObj" ] }`,
			[]testUserType{obj, arr, refToObj},
			[]ElementType{ElementTypeScalar, ElementTypeScalar, ElementTypeObject, ElementTypeArray, ElementTypeObject},
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

	elements := Dereference(jSchema)

	actual := make([]ElementType, 0, len(elements))
	for _, e := range elements {
		actual = append(actual, e.Type())
	}

	require.Equal(t, data.expected, actual)
}
