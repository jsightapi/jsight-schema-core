package ischema

import (
	"testing"

	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/lexeme"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"
	"github.com/jsightapi/jsight-schema-core/test"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[lexeme.LexEventType]Node{
			lexeme.LiteralBegin:    &LiteralNode{},
			lexeme.ObjectBegin:     &ObjectNode{},
			lexeme.ArrayBegin:      &ArrayNode{},
			lexeme.MixedValueBegin: &MixedValueNode{},
		}

		for lexType, expected := range cc {
			t.Run(lexType.String(), func(t *testing.T) {
				actual := NewNode(lexeme.NewLexEvent(lexType, 0, 0, nil))
				assert.IsType(t, expected, actual)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		test.PanicsWithErr(t, errs.ErrRuntimeFailure.F(), func() {
			NewNode(lexeme.NewLexEvent(lexeme.NewLine, 0, 0, nil))
		})
	})
}

func TestIsOptionalNode(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]struct {
			given    func(*testing.T) Node
			expected bool
		}{
			"node without optional constraints": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.OptionalConstraintType).
						Return(nil)
					return n
				},
				false,
			},
			"not a bool keeper": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.OptionalConstraintType).
						Return(constraint.Max{})
					return n
				},
				false,
			},
			"false": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.OptionalConstraintType).
						Return(constraint.NewOptional(bytes.NewBytes("false")))
					return n
				},
				false,
			},
			"true": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.OptionalConstraintType).
						Return(constraint.NewOptional(bytes.NewBytes("true")))
					return n
				},
				true,
			},
		}

		for n, c := range cc {
			t.Run(n, func(t *testing.T) {
				actual := IsOptionalNode(c.given(t))
				assert.Equal(t, c.expected, actual)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.Panics(t, func() {
			IsOptionalNode(nil)
		})
	})
}

func TestIsNullableNode(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]struct {
			given    func(*testing.T) Node
			expected bool
		}{
			"node without nullable constraints": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.NullableConstraintType).
						Return(nil)
					return n
				},
				false,
			},
			"not a bool keeper": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.NullableConstraintType).
						Return(constraint.Max{})
					return n
				},
				false,
			},
			"false": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.NullableConstraintType).
						Return(constraint.NewNullable(bytes.NewBytes("false")))
					return n
				},
				false,
			},
			"true": {
				func(t *testing.T) Node {
					n := NewMockNode(t)
					n.
						On("Constraint", constraint.NullableConstraintType).
						Return(constraint.NewNullable(bytes.NewBytes("true")))
					return n
				},
				true,
			},
		}

		for n, c := range cc {
			t.Run(n, func(t *testing.T) {
				actual := IsNullableNode(c.given(t))
				assert.Equal(t, c.expected, actual)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.Panics(t, func() {
			IsNullableNode(nil)
		})
	})
}
