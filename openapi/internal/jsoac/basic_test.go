package jsoac

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type testConverterData struct {
	jsight  string
	openapi string
}

type testComplexConverterData struct {
	jsight    string
	openapi   string
	userTypes []testUserType
}

type testUserType struct {
	name   string
	jsight string
}

var catUserType = testUserType{
	"@cat",
	`{ "catName": "Tom" }`,
}

var dogUserType = testUserType{
	"@dog",
	`{ "dogName": "Max" }`,
}

var catEmailUserType = testUserType{
	"@catEmail",
	`"abc@cat.com" // { regex: "^[a-z]+@cat.com$" }`,
}

var dogEmailUserType = testUserType{
	"@dogEmail",
	`"abc@dog.com" // { regex: "^[a-z]+@dog.com$" }`,
}

var stringIDUserType = testUserType{
	"@stringId",
	`"abc-123"`,
}

var integerIDUserType = testUserType{
	"@integerId",
	`123`,
}

func (t testConverterData) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

func (t testComplexConverterData) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

func assertJSightToOpenAPIConverter(t *testing.T, data testConverterData) {
	d := testComplexConverterData{
		jsight:    data.jsight,
		openapi:   data.openapi,
		userTypes: []testUserType{},
	}
	assertJSightToOpenAPIComplexConverter(t, d)
}

func buildJSchema(t *testing.T, jsight string, userTypes []testUserType) *jschema.JSchema {
	jSchema := jschema.New("root", jsight)

	for _, ut := range userTypes {
		err := jSchema.AddType(ut.name, jschema.New(ut.name, ut.jsight))
		require.NoError(t, err)
	}

	err := jSchema.Check()
	require.NoError(t, err)

	return jSchema
}

func assertJSightToOpenAPIComplexConverter(t *testing.T, data testComplexConverterData) {
	jSchema := buildJSchema(t, data.jsight, data.userTypes)

	o := New(jSchema)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	jsonString := string(json)
	require.JSONEq(t, data.openapi, jsonString, "Actual: "+jsonString)
}
