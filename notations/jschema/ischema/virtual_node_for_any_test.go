package ischema

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"
)

func Test_VirtualNodeForAny_Data(t *testing.T) {
	n := VirtualNodeForAny()

	c := n.Constraint(constraint.TypeConstraintType)
	require.NotNil(t, c)
	require.Equal(t, `"any"`, c.(*constraint.TypeConstraint).Bytes().String())
}

func Test_VirtualNodeForAny_Same(t *testing.T) {
	n1 := VirtualNodeForAny()
	n2 := VirtualNodeForAny()

	assert.Same(t, n1, n2)
}
