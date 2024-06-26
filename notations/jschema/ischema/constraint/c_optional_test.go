package constraint

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestNewOptional(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]bool{
			"true":  true,
			"false": false,
		}

		for given, expected := range cc {
			t.Run(given, func(t *testing.T) {
				c := NewOptional(bytes.NewBytes(given))
				assert.Equal(t, expected, c.value)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `Invalid value in the "optional" rule. Learn about the rules here: https://jsight.io/docs/jsight-schema-0-3#rules`, func() {
			NewOptional(bytes.NewBytes("foo"))
		})
	})
}

func TestOptional_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(t, Optional{}, allJSONTypes...)
}

func TestOptional_Type(t *testing.T) {
	assert.Equal(t, OptionalConstraintType, NewOptional(bytes.NewBytes("true")).Type())
}

func TestOptional_String(t *testing.T) {
	cc := map[string]string{
		"false": "[ UNVERIFIABLE CONSTRAINT ] optional: false",
		"true":  "[ UNVERIFIABLE CONSTRAINT ] optional: true",
	}

	for given, expected := range cc {
		t.Run(given, func(t *testing.T) {
			assert.Equal(t, expected, NewOptional(bytes.NewBytes(given)).String())
		})
	}
}

func TestOptional_Bool(t *testing.T) {
	cc := map[string]bool{
		"false": false,
		"true":  true,
	}

	for given, expected := range cc {
		t.Run(given, func(t *testing.T) {
			assert.Equal(t, expected, NewOptional(bytes.NewBytes(given)).Bool())
		})
	}
}

func TestOptional_ASTNode(t *testing.T) {
	cc := []bool{true, false}

	for _, c := range cc {
		t.Run(strconv.FormatBool(c), func(t *testing.T) {
			assert.Equal(t, schema.RuleASTNode{
				TokenType:  schema.TokenTypeBoolean,
				Value:      strconv.FormatBool(c),
				Properties: &schema.RuleASTNodes{},
				Source:     schema.RuleASTNodeSourceManual,
			}, Optional{value: c}.ASTNode())
		})
	}
}
