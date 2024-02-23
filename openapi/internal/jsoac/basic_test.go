package jsoac

import (
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

var testCatUserType = testUserType{
	"@cat",
	`{ "catName": "Tom" }`,
}

var testDogUserType = testUserType{
	"@dog",
	`{ "dogName": "Max" }`,
}

func assertJSightToOpenAPIConverter(t *testing.T, data testConverterData) {
	d := testComplexConverterData{
		jsight:    data.jsight,
		openapi:   data.openapi,
		userTypes: []testUserType{},
	}
	assertJSightToOpenAPIComplexConverter(t, d)
}

func assertJSightToOpenAPIComplexConverter(t *testing.T, data testComplexConverterData) {
	j := jschema.New("root", data.jsight)

	for _, ut := range data.userTypes {
		err := j.AddType(ut.name, jschema.New(ut.name, ut.jsight))
		require.NoError(t, err)
	}

	err := j.Check()
	require.NoError(t, err)

	o := New(j)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	jsonString := string(json)
	require.JSONEq(t, data.openapi, jsonString, "Actual: "+jsonString)
}
