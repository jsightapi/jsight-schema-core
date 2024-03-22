package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/regex"

	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

func Test_NewJSchemaObject(t *testing.T) {
	j := jschema.New("TestSchemaName", `{}`)
	err := j.Check()
	require.NoError(t, err)

	t.Run("from JSchema", func(t *testing.T) {
		o := NewSchemaObject(j)
		o.SetDescription(`Any description string & "quoted string" & \*\/ \*\/`)

		json, err := o.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"type": "object", "properties": {}, "additionalProperties": false, "description": "Any description string & \"quoted string\" & \\*\\/ \\*\\/"}`, string(json))
	})
}

func Test_NewRegexObject(t *testing.T) {
	r := regex.New("TestSchemaName", `/OK/`)
	err := r.Check()
	require.NoError(t, err)

	t.Run("from RSchema", func(t *testing.T) {
		o := NewSchemaObject(r)
		o.SetDescription(`Any description string & "quoted string" & \*\/ \*\/`)

		json, err := o.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"type": "string", "pattern": "OK"}`, string(json))
	})
}
