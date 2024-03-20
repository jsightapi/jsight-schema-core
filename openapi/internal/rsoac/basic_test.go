package rsoac

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

func (t testConverterData) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

func assertJSightToOpenAPIConverter(t *testing.T, data testConverterData) {
	d := testComplexConverterData{
		jsight:    data.jsight,
		openapi:   data.openapi
	}
	assertJSightToOpenAPIComplexConverter(t, d)
}

func buildRSchema(t *testing.T, jsight string, userTypes []testUserType) *jschema.JSchema {
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
	rSchema := buildRSchema(t, data.jsight)

	o := New(rSchema)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	jsonString := string(json)
	require.JSONEq(t, data.openapi, jsonString, "Actual: "+jsonString)
}
