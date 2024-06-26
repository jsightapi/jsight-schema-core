package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/json"

	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestNewMaxLength(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cnstr := NewMaxLength(bytes.NewBytes("10"))

		assert.EqualValues(t, 10, cnstr.value)
	})

	t.Run("negative", func(t *testing.T) {
		ss := []string{
			"not a number",
			"3.14",
			"-12",
		}

		for _, s := range ss {
			t.Run(s, func(t *testing.T) {
				assert.PanicsWithError(t, `Invalid value in the "maxLength" rule. Learn about the rules here: https://jsight.io/docs/jsight-schema-0-3#rules`, func() {
					NewMaxLength(bytes.NewBytes(s))
				})
			})
		}
	})
}

func TestMaxLength_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(t, MaxLength{}, json.TypeString)
}

func TestMaxLength_Type(t *testing.T) {
	assert.Equal(t, MaxLengthConstraintType, NewMaxLength(bytes.NewBytes("1")).Type())
}

func TestMaxLength_String(t *testing.T) {
	assert.Equal(t, "maxLength: 1", NewMaxLength(bytes.NewBytes("1")).String())
}

func TestMaxLength_Validate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := []string{
			"",
			"foo",
			"0123456789",
		}

		for _, given := range cc {
			t.Run(given, func(t *testing.T) {
				assert.NotPanics(t, func() {
					NewMaxLength(bytes.NewBytes("10")).Validate(bytes.NewBytes(given))
				})
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, "The length of the string in the example violates the rule `\"maxLength\": \"10\"`", func() {
			NewMaxLength(bytes.NewBytes("10")).Validate(bytes.NewBytes("0123456789A"))
		})
	})
}

func TestMaxLength_ASTNode(t *testing.T) {
	assert.Equal(t, schema.RuleASTNode{
		TokenType:  schema.TokenTypeNumber,
		Value:      "1",
		Properties: &schema.RuleASTNodes{},
		Source:     schema.RuleASTNodeSourceManual,
	}, NewMaxLength(bytes.NewBytes("1")).ASTNode())
}

func TestMaxLength_Value(t *testing.T) {
	assert.Equal(t, uint(1), NewMaxLength(bytes.NewBytes("1")).Value())
}
