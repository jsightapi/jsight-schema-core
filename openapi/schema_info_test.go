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

	p := info.PropertiesInfos()

	p.Next()
	k := p.GetKey()
	v := p.GetInfo()
	assert.Equal(t, "k1", k)
	assert.Equal(t, "property annotation 1", v.Annotation())
	assert.True(t, v.Optional())

	expJson := `
		{
			"type": "string",
			"optional": true,
			"example": "v1",
			"description": "property annotation 1"
		}
	`
	actualJson, err := v.SchemaObject().MarshalJSON()
	require.NoError(t, err)
	assert.JSONEq(t, expJson, string(actualJson))

	p.Next()
	k = p.GetKey()
	v = p.GetInfo()
	assert.Equal(t, "k2", k)
	assert.Equal(t, "property annotation 2", v.Annotation())
	assert.False(t, v.Optional())

	p.Next()
	k = p.GetKey()
	v = p.GetInfo()
	assert.Equal(t, "k3", k)
	assert.Equal(t, "property annotation 3", v.Annotation())
	assert.False(t, v.Optional())
}
