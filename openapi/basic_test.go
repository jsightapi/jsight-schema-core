package openapi

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type testUserType struct {
	name   string
	jsight string
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
