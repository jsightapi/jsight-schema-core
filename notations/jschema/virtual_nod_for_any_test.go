package jschema

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"
)

func Test_VirtualNodeForAny_Data(t *testing.T) {
	n := VirtualNodeForAny()

	assert.NotNil(t, n.Constraint(constraint.AnyConstraintType))
}

func Test_VirtualNodeForAny_Same(t *testing.T) {
	n1 := VirtualNodeForAny()
	n2 := VirtualNodeForAny()

	assert.Same(t, n1, n2)
}
