package ischema

import (
	"testing"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/jsightapi/jsight-schema-core/json"
	"github.com/jsightapi/jsight-schema-core/lexeme"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"
	"github.com/jsightapi/jsight-schema-core/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMixedValueNode(t *testing.T) {
	e := lexeme.NewLexEvent(lexeme.MixedValueBegin, 0, 0, nil)
	n := NewMixedValueNode(e)

	assert.Nil(t, n.parent)
	assert.Equal(t, json.TypeMixed, n.jsonType)
	assert.Equal(t, e, n.schemaLexEvent)
	assert.Equal(t, &Constraints{}, n.constraints)
}

func TestMixedValueNode_AddConstraint(t *testing.T) {
	t.Run("Type constraint", func(t *testing.T) {
		n := createFakeMixedValueNode()
		n.AddConstraint(createFakeTypeConstraint("@foo"))

		assert.Equal(t, []string{"@foo"}, n.types)
	})

	t.Run("Or constraint", func(t *testing.T) {
		n := createFakeMixedValueNode()
		n.AddConstraint(constraint.NewOr(schema.RuleASTNodeSourceManual))

		assert.Equal(t, []string(nil), n.types)
	})

	t.Run("TypeList constraint", func(t *testing.T) {
		c := constraint.NewTypesList(schema.RuleASTNodeSourceManual)
		c.AddName("@foo", "@foo", schema.RuleASTNodeSourceManual)
		c.AddName("@bar", "@bar", schema.RuleASTNodeSourceManual)

		n := createFakeMixedValueNode()
		n.AddConstraint(c)

		assert.Equal(t, []string{"@foo", "@bar"}, n.types)
	})
}

func TestMixedValueNode_addTypeConstraint(t *testing.T) {
	t.Run("not exists", func(t *testing.T) {
		const name = "@foo"
		n := createFakeMixedValueNode()
		n.addTypeConstraint(createFakeTypeConstraint(name))

		c, ok := n.baseNode.constraints.Get(constraint.TypeConstraintType)
		require.True(t, ok)
		require.IsType(t, &constraint.TypeConstraint{}, c)

		assert.Equal(t, name, c.(*constraint.TypeConstraint).Bytes().String())
	})

	t.Run("exists", func(t *testing.T) {
		cc := map[string]struct {
			exists             *constraint.TypeConstraint
			new                *constraint.TypeConstraint
			expected           *constraint.TypeConstraint
			expectedSchemaType string
		}{
			"equal, not mixed": {
				createFakeTypeConstraint("@foo"),
				createFakeTypeConstraint("@foo"),
				createFakeTypeConstraint("@foo"),
				"@foo",
			},
			"equal, mixed": {
				createFakeTypeConstraint("mixed"),
				createFakeTypeConstraint("mixed"),
				createFakeTypeConstraint("mixed"),
				"mixed",
			},
			"not equal, new is mixed": {
				createFakeTypeConstraint("@foo"),
				createFakeTypeConstraint("mixed"),
				createFakeTypeConstraint("mixed"),
				"mixed",
			},
		}

		for n, c := range cc {
			t.Run(n, func(t *testing.T) {
				n := createFakeMixedValueNode()
				n.schemaType = "should be changed"
				n.baseNode.AddConstraint(c.exists)

				n.addTypeConstraint(c.new)

				actual := n.baseNode.constraints.GetValue(constraint.TypeConstraintType)
				assert.Equal(t, c.expected, actual)
			})
		}

		t.Run("not equal, new isn't mixed", func(t *testing.T) {
			assert.PanicsWithError(t, `Duplicate rule "type"`, func() {
				n := createFakeMixedValueNode()
				n.schemaType = "should be changed"
				n.baseNode.AddConstraint(createFakeTypeConstraint("@foo"))

				n.addTypeConstraint(createFakeTypeConstraint("@bar"))
			})
		})
	})
}

func Test_addOrConstraint(t *testing.T) {
	t.Run("without type constraint", func(t *testing.T) {
		expected := constraint.NewOr(schema.RuleASTNodeSourceManual)

		n := createFakeMixedValueNode()
		n.addOrConstraint(expected)

		actual := n.baseNode.constraints.GetValue(constraint.OrConstraintType)
		assert.Equal(t, expected, actual)
	})

	t.Run("with type constraint", func(t *testing.T) {
		expected := constraint.NewOr(schema.RuleASTNodeSourceManual)

		n := createFakeMixedValueNode()
		n.baseNode.AddConstraint(createFakeTypeConstraint("@foo"))

		n.addOrConstraint(expected)

		actual := n.baseNode.constraints.GetValue(constraint.OrConstraintType)
		assert.Equal(t, expected, actual)

		actual = n.baseNode.constraints.GetValue(constraint.TypeConstraintType)
		assert.Equal(t, createFakeTypeConstraint(`"mixed"`), actual)
	})
}

func TestMixedValueNode_Grow(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		n := createFakeMixedValueNode()
		n.parent = NewObjectNode(lexeme.NewLexEvent(lexeme.ObjectBegin, 0, 0, nil))

		cc := map[lexeme.LexEventType]Node{
			lexeme.MixedValueBegin: n,
			lexeme.MixedValueEnd:   n.parent,
		}

		for lexType, expected := range cc {
			t.Run(lexType.String(), func(t *testing.T) {
				actual, ok := n.Grow(lexeme.NewLexEvent(lexType, 0, 0, fs.NewFile("", "foo")))
				assert.Equal(t, expected, actual)
				assert.False(t, ok)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		test.PanicsWithErr(t,
			errs.ErrUnexpectedLexicalEvent.F(lexeme.ObjectBegin.String(), "in mixed value node"),
			func() {
				createFakeMixedValueNode().
					Grow(lexeme.NewLexEvent(lexeme.ObjectBegin, 0, 0, nil))
			},
		)
	})
}

func createFakeMixedValueNode() *MixedValueNode {
	return NewMixedValueNode(lexeme.NewLexEvent(lexeme.MixedValueBegin, 0, 0, nil))
}

func createFakeTypeConstraint(name string) *constraint.TypeConstraint {
	return constraint.NewType(bytes.NewBytes(name), schema.RuleASTNodeSourceManual)
}
