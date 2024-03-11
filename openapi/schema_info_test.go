package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

func Test_SchemaInfo(t *testing.T) {
	j := jschema.New("root", `{ // object annotation
"k1": "v1", // {optional: true} - property annotation 1
"k2": "v2", // {optional: false} - property annotation 2
"k3": "v3" // property annotation 3
}`)
	err := j.Check()
	require.NoError(t, err)

	info := NewSchemaInfo(j)

	assert.Equal(t, "object annotation", info.Annotation())
	assert.False(t, info.Optional())

	props := info.PropertiesInfos()

	b := props.Next()
	require.True(t, b)

	k := props.GetKey()
	v := props.GetInfo()
	assert.Equal(t, "k1", k)
	assert.Equal(t, "property annotation 1", v.Annotation())
	assert.True(t, v.Optional())
	assertPropertyJson(t, v.SchemaObject(), `{
		"type": "string",
		"example": "v1",
		"description": "property annotation 1"
	}`)

	b = props.Next()
	require.True(t, b)

	k = props.GetKey()
	v = props.GetInfo()
	assert.Equal(t, "k2", k)
	assert.Equal(t, "property annotation 2", v.Annotation())
	assert.False(t, v.Optional())
	assertPropertyJson(t, v.SchemaObject(), `{
		"type": "string",
		"example": "v2",
		"description": "property annotation 2"
	}`)

	b = props.Next()
	require.True(t, b)

	k = props.GetKey()
	v = props.GetInfo()
	assert.Equal(t, "k3", k)
	assert.Equal(t, "property annotation 3", v.Annotation())
	assert.False(t, v.Optional())
	assertPropertyJson(t, v.SchemaObject(), `{
		"type": "string",
		"example": "v3",
		"description": "property annotation 3"
	}`)

	b = props.Next()
	require.False(t, b)
}

func assertPropertyJson(t *testing.T, so SchemaObject, expectedJSON string) {
	actualJson, err := so.MarshalJSON()
	require.NoError(t, err)
	assert.JSONEq(t, expectedJSON, string(actualJson))
}

/*
JSIGHT 0.3

GET /

	Request
		Headers
			@ref1
	Body any

TYPE @ref1

	@ref2

TYPE @ref2

	{
		"k1": "v1", // property annotation 1
		"k2": "v2" // {optional: true} - property annotation 2
	}
*/
func Test_SchemaInfo_PropertiesInfos(t *testing.T) {
	j := jschema.New("root", "@ref1")

	err := j.AddType("@ref1", jschema.New("@ref1", "@ref2"))
	require.NoError(t, err)

	err = j.AddType("@ref2", jschema.New("@ref2", `{
		"k1": "v1", // property annotation 1
		"k2": "v2" // {optional: true} - property annotation 2
	}`))
	require.NoError(t, err)

	err = j.Check()
	require.NoError(t, err)

	info := NewSchemaInfo(j)
	props := info.PropertiesInfos()

	b := props.Next()
	require.True(t, b)

	k := props.GetKey()
	v := props.GetInfo()
	assert.Equal(t, "k1", k)
	assert.Equal(t, "property annotation 1", v.Annotation())
	assert.False(t, v.Optional())
	assertPropertyJson(t, v.SchemaObject(), `{
		"type": "string",
		"example": "v1",
		"description": "property annotation 1"
	}`)

	b = props.Next()
	require.True(t, b)

	k = props.GetKey()
	v = props.GetInfo()
	assert.Equal(t, "k2", k)
	assert.Equal(t, "property annotation 2", v.Annotation())
	assert.True(t, v.Optional())
	assertPropertyJson(t, v.SchemaObject(), `{
		"type": "string",
		"example": "v2",
		"description": "property annotation 2"
	}`)

	b = props.Next()
	require.False(t, b)
}
