package openapi

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

func TestNewSchemaObject(t *testing.T) {
	j := jschema.New("TestSchemaName", `{}`)
	err := j.Check()
	require.NoError(t, err)

	t.Run("positive", func(t *testing.T) {
		t.Run("from JSchema", func(t *testing.T) {
			o := NewSchemaObject(j)
			json, err := o.MarshalJSON()
			require.NoError(t, err)
			require.JSONEq(t, `{"type": "object", "properties": {}}`, string(json))
		})

		// TODO regex
	})
}
