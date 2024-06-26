package loader

import (
	"fmt"
	"testing"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/lexeme"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/internal/mocks"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"
	"github.com/jsightapi/jsight-schema-core/test"

	"github.com/stretchr/testify/assert"
)

func TestSchemaCompiler_checkMinAndMax(t *testing.T) {
	cc := map[string]struct {
		node        func(*testing.T) ischema.Node
		expectedErr string
	}{
		"nil min, nil max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxConstraintType).Return(nil)
				return m
			},
		},
		"nil min, not nil max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("42")))
				return m
			},
		},
		"nil min, not nil max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				max := constraint.NewMax(bytes.NewBytes("42"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
		},
		"not nil min, nil max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("42")))
				m.On("Constraint", constraint.MaxConstraintType).Return(nil)
				return m
			},
		},
		"not nil min (exclusive), nil max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("42"))
				min.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(nil)
				return m
			},
		},
		"min < max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("1")))
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("2")))
				return m
			},
		},
		"min (exclusive) < max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("1"))
				min.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("2")))
				return m
			},
		},
		"min < max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				max := constraint.NewMax(bytes.NewBytes("2"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("1")))
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
		},
		"min (exclusive) < max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("1"))
				min.SetExclusive(true)
				max := constraint.NewMax(bytes.NewBytes("2"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
		},
		"min = max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("1")))
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("1")))
				return m
			},
		},
		"min (exclusive) = max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("1"))
				min.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("1")))
				return m
			},
			expectedErr: `The value of the rule "min" should be less than the value of the rule "max"`,
		},
		"min = max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				max := constraint.NewMax(bytes.NewBytes("1"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("1")))
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
			expectedErr: `The value of the rule "min" should be less than the value of the rule "max"`,
		},
		"min (exclusive) = max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("1"))
				min.SetExclusive(true)
				max := constraint.NewMax(bytes.NewBytes("1"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
			expectedErr: `The value of the rule "min" should be less than the value of the rule "max"`,
		},
		"min > max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("2")))
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("1")))
				return m
			},
			expectedErr: `The value of the rule "min" should be less or equal to the value of the rule "max"`,
		},
		"min (exclusive) > max": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("2"))
				min.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(constraint.NewMax(bytes.NewBytes("1")))
				return m
			},
			expectedErr: `The value of the rule "min" should be less than the value of the rule "max"`,
		},
		"min > max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				max := constraint.NewMax(bytes.NewBytes("1"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(constraint.NewMin(bytes.NewBytes("2")))
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
			expectedErr: `The value of the rule "min" should be less than the value of the rule "max"`,
		},
		"min (exclusive) > max (exclusive)": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				min := constraint.NewMin(bytes.NewBytes("2"))
				min.SetExclusive(true)
				max := constraint.NewMax(bytes.NewBytes("1"))
				max.SetExclusive(true)
				m.On("Constraint", constraint.MinConstraintType).Return(min)
				m.On("Constraint", constraint.MaxConstraintType).Return(max)
				return m
			},
			expectedErr: `The value of the rule "min" should be less than the value of the rule "max"`,
		},
	}

	for n, c := range cc {
		t.Run(n, func(t *testing.T) {
			err := schemaCompiler{}.checkMinAndMax(c.node(t))
			if c.expectedErr != "" {
				assert.EqualError(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSchemaCompiler_checkMinLengthAndMaxLength(t *testing.T) {
	cc := map[string]struct {
		node        func(*testing.T) ischema.Node
		expectedErr string
	}{
		"nil minLength, nil maxLength": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinLengthConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxLengthConstraintType).Return(nil)
				return m
			},
		},
		"nil minLength, not nil maxLength": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinLengthConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxLengthConstraintType).Return(constraint.NewMaxLength(bytes.NewBytes("42")))
				return m
			},
		},
		"not nil minLength, nil maxLength": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinLengthConstraintType).Return(constraint.NewMinLength(bytes.NewBytes("42")))
				m.On("Constraint", constraint.MaxLengthConstraintType).Return(nil)
				return m
			},
		},
		"minLength < maxLength": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinLengthConstraintType).Return(constraint.NewMinLength(bytes.NewBytes("1")))
				m.On("Constraint", constraint.MaxLengthConstraintType).Return(constraint.NewMaxLength(bytes.NewBytes("2")))
				return m
			},
		},
		"minLength = maxLength": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinLengthConstraintType).Return(constraint.NewMinLength(bytes.NewBytes("2")))
				m.On("Constraint", constraint.MaxLengthConstraintType).Return(constraint.NewMaxLength(bytes.NewBytes("2")))
				return m
			},
		},
		"minLength > maxLength": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinLengthConstraintType).Return(constraint.NewMinLength(bytes.NewBytes("2")))
				m.On("Constraint", constraint.MaxLengthConstraintType).Return(constraint.NewMaxLength(bytes.NewBytes("1")))
				return m
			},
			expectedErr: `The value of the rule "minLength" should be less or equal to the value of the rule "maxLength"`,
		},
	}

	for n, c := range cc {
		t.Run(n, func(t *testing.T) {
			err := schemaCompiler{}.checkMinLengthAndMaxLength(c.node(t))
			if c.expectedErr != "" {
				assert.EqualError(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSchemaCompiler_checkMinItemsAndMaxItems(t *testing.T) {
	cc := map[string]struct {
		node        func(*testing.T) ischema.Node
		expectedErr string
	}{
		"nil minItems, nil maxItems": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinItemsConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxItemsConstraintType).Return(nil)
				return m
			},
		},
		"nil minItems, not nil maxItems": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinItemsConstraintType).Return(nil)
				m.On("Constraint", constraint.MaxItemsConstraintType).Return(constraint.NewMaxItems(bytes.NewBytes("42")))
				return m
			},
		},
		"not nil minItems, nil maxItems": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinItemsConstraintType).Return(constraint.NewMinItems(bytes.NewBytes("42")))
				m.On("Constraint", constraint.MaxItemsConstraintType).Return(nil)
				return m
			},
		},
		"minItems < maxItems": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinItemsConstraintType).Return(constraint.NewMinItems(bytes.NewBytes("1")))
				m.On("Constraint", constraint.MaxItemsConstraintType).Return(constraint.NewMaxItems(bytes.NewBytes("2")))
				return m
			},
		},
		"minItems = maxItems": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinItemsConstraintType).Return(constraint.NewMinItems(bytes.NewBytes("2")))
				m.On("Constraint", constraint.MaxItemsConstraintType).Return(constraint.NewMaxItems(bytes.NewBytes("2")))
				return m
			},
		},
		"minItems > maxItems": {
			node: func(t *testing.T) ischema.Node {
				m := mocks.NewNode(t)
				m.On("Constraint", constraint.MinItemsConstraintType).Return(constraint.NewMinItems(bytes.NewBytes("2")))
				m.On("Constraint", constraint.MaxItemsConstraintType).Return(constraint.NewMaxItems(bytes.NewBytes("1")))
				return m
			},
			expectedErr: `The value of the rule "minItems" should be less or equal to the value of the rule "maxItems"`,
		},
	}

	for n, c := range cc {
		t.Run(n, func(t *testing.T) {
			err := schemaCompiler{}.checkMinItemsAndMaxItems(c.node(t))
			if c.expectedErr != "" {
				assert.EqualError(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSchemaCompiler_precisionConstraint(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]func(*mocks.Node){
			"without constraints": func(n *mocks.Node) {
				n.On("Constraint", constraint.PrecisionConstraintType).Return(nil)
			},
			"with only precision constraint": func(n *mocks.Node) {
				n.On("Constraint", constraint.PrecisionConstraintType).Return(constraint.Precision{})
				n.On("Constraint", constraint.TypeConstraintType).Return(nil)
			},
			"type is decimal": func(n *mocks.Node) {
				n.On("Constraint", constraint.PrecisionConstraintType).Return(constraint.Precision{})
				n.On("Constraint", constraint.TypeConstraintType).Return(constraint.NewType(
					bytes.NewBytes("decimal"),
					schema.RuleASTNodeSourceManual,
				))
			},
		}

		for name, fn := range cc {
			t.Run(name, func(t *testing.T) {
				n := mocks.NewNode(t)
				fn(n)
				schemaCompiler{}.precisionConstraint(n)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `The rule "precision" is not compatible with the "foo" type. Learn more about the rules and types compatibility here: https://jsight.io/docs/jsight-schema-0-3#appendix-1-a-table-of-all-built-in-types-and-rules`, func() {
			n := mocks.NewNode(t)
			n.On("Constraint", constraint.PrecisionConstraintType).Return(constraint.Precision{})
			n.On("Constraint", constraint.TypeConstraintType).Return(constraint.NewType(
				bytes.NewBytes("foo"),
				schema.RuleASTNodeSourceManual,
			))
			schemaCompiler{}.precisionConstraint(n)
		})
	})
}

func TestSchemaCompiler_emptyArray(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := []ischema.Node{
			nil,
			&ischema.ObjectNode{},
			func() *ischema.ArrayNode {
				n := &ischema.ArrayNode{}
				n.Grow(newFakeLexEvent(lexeme.ArrayItemBegin))
				n.Grow(newFakeLexEventWithValue(lexeme.LiteralBegin, "foo"))
				return n
			}(),
			&ischema.ArrayNode{},
			func() ischema.Node {
				n := ischema.NewNode(newFakeLexEvent(lexeme.ArrayBegin))
				n.AddConstraint(constraint.NewMinItems(bytes.NewBytes("0")))
				n.AddConstraint(constraint.NewMaxItems(bytes.NewBytes("0")))
				return n
			}(),
		}

		for i, given := range cc {
			t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
				assert.NotPanics(t, func() {
					schemaCompiler{}.emptyArray(given)
				})
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		cc := map[string]ischema.Node{
			"min": func() ischema.Node {
				n := ischema.NewNode(newFakeLexEvent(lexeme.ArrayBegin))
				n.AddConstraint(constraint.NewMinItems(bytes.NewBytes("1")))
				return n
			}(),
			"max": func() ischema.Node {
				n := ischema.NewNode(newFakeLexEvent(lexeme.ArrayBegin))
				n.AddConstraint(constraint.NewMaxItems(bytes.NewBytes("1")))
				return n
			}(),
		}

		for n, given := range cc {
			t.Run(n, func(t *testing.T) {
				test.PanicsWithErr(t, errs.ErrIncorrectConstraintValueForEmptyArray.F(), func() {
					schemaCompiler{}.emptyArray(given)
				})
			})
		}
	})
}

func BenchmarkSchemaCompiler_emptyArray(b *testing.B) {
	n := ischema.NewNode(newFakeLexEvent(lexeme.ArrayBegin))
	n.AddConstraint(constraint.NewMinItems(bytes.NewBytes("0")))
	n.AddConstraint(constraint.NewMaxItems(bytes.NewBytes("0")))

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		schemaCompiler{}.emptyArray(n)
	}
}
