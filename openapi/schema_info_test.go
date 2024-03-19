package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/regex"
)

func Test_NewRSchemaInfo(t *testing.T) {
	rSchema := &regex.RSchema{}

	info := NewRSchemaInfo(rSchema)

	assert.Equal(t, SchemaInfoTypeRegex, info.Type())
	assert.Equal(t, "", info.Annotation())

	// TODO after rsoac.go
	// so := info.SchemaObject()
	// so.SetDescription("Some text 2")
	// json, err := so.MarshalJSON()
	//
	// require.NoError(t, err)

	// jsonString := string(json)
	// require.JSONEq(t, `...`, jsonString, "Actual: "+jsonString)
}

func Test_NewJSchemaInfo(t *testing.T) {
	jSchema := buildJSchema(t, `123 // {min: 1} - Some text`, []testUserType{})

	info := NewJSchemaInfo(jSchema)

	assert.Equal(t, SchemaInfoTypeScalar, info.Type())
	assert.Equal(t, "Some text", info.Annotation())

	so := info.SchemaObject()
	so.SetDescription("Some text 2")
	json, err := so.MarshalJSON()

	require.NoError(t, err)

	jsonString := string(json)
	require.JSONEq(t, `{"type":"integer","example":123,"minimum":1,"description":"Some text 2"}`, jsonString, "Actual: "+jsonString)
}
