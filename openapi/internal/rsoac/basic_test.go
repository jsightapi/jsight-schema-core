package rsoac

import (
	"github.com/jsightapi/jsight-schema-core/notations/regex"

	"github.com/stretchr/testify/require"

	"regexp"

	"testing"
)

type testConverterRegex struct {
	jsight  string
	openapi string
}

func (t testConverterRegex) name() string {
	re := regexp.MustCompile(`[\s/]`)
	return re.ReplaceAllString(t.jsight, "_")
}

func assertJSightToOpenAPIConverter(t *testing.T, data testConverterRegex) {
	rSchema := buildRSchema(t, data.jsight)

	o := New(rSchema)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	jsonString := string(json)
	require.JSONEq(t, data.openapi, jsonString, "Actual: "+jsonString)
}

func buildRSchema(t *testing.T, jsight string) *regex.RSchema {
	rSchema := regex.New("root", jsight)

	err := rSchema.Check()
	require.NoError(t, err)

	return rSchema
}
