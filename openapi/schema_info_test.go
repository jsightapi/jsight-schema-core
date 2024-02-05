package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

func TestObjectProperties(t *testing.T) {
	// JSIGHT 0.3
	// TYPE @userType2
	//   @userType1

	ut1 := jschema.New("@userType1", `{ // object annotation
"k": "v" // {optional: true} - property annotation
}`)
	err := ut1.Check()
	require.NoError(t, err)

	// TODO ut2 := jschema.New("@userType2", `@userType1`)
	//err = ut2.AddType("@userType1", ut1)
	//require.NoError(t, err)
	//err = ut2.Check()
	//require.NoError(t, err)

	info := NewSchemaInfo(ut1)

	assert.Equal(t, "object annotation", info.Annotation())
	assert.False(t, info.Optional())

	for _, child := range info.NestedObjectProperties() {
		assert.Equal(t, "property annotation", child.Annotation())
		assert.True(t, child.Optional())
	}
}
