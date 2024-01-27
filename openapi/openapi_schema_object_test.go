package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSchemaObject(t *testing.T) {
	j := jschema.New("TestSchemaName", `{}`)
	err := j.Check()
	require.NoError(t, err)

	o := NewSchemaObject(j)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	require.JSONEq(t, `{"type": "object", "properties": {}}`, string(json))
}
