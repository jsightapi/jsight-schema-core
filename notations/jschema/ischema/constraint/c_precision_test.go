package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/json"

	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestNewPrecision(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		c := NewPrecision(bytes.NewBytes("10"))
		assert.Equal(t, uint(10), c.value)
	})

	t.Run("negative", func(t *testing.T) {
		cc := map[string]string{
			"-10":  `Invalid value in the "precision" rule. Learn about the rules here: https://jsight.io/docs/jsight-schema-0-3#rules`,
			"0":    "Precision can not be zero",
			"3.14": `Invalid value in the "precision" rule. Learn about the rules here: https://jsight.io/docs/jsight-schema-0-3#rules`,
		}

		for given, expected := range cc {
			t.Run(given, func(t *testing.T) {
				assert.PanicsWithError(t, expected, func() {
					NewPrecision(bytes.NewBytes(given))
				})
			})
		}
	})
}

func TestPrecision_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(t, Precision{}, json.TypeFloat)
}

func TestPrecision_Type(t *testing.T) {
	assert.Equal(t, PrecisionConstraintType, NewPrecision(bytes.NewBytes("1")).Type())
}

func TestPrecision_String(t *testing.T) {
	assert.Equal(t, "precision: 1", NewPrecision(bytes.NewBytes("1")).String())
}

func TestPrecision_Validate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]string{
			"3.14":  "",
			"3.1":   "",
			"3":     "",
			"3.142": "The value in the example violates the rule `\"precision\": 2` (exclusive)",
		}

		for given, expectedError := range cc {
			t.Run(given, func(t *testing.T) {
				cnstr := NewPrecision(bytes.NewBytes("2"))
				if expectedError != "" {
					assert.PanicsWithError(t, expectedError, func() {
						cnstr.Validate(bytes.NewBytes(given))
					})
				} else {
					assert.NotPanics(t, func() {
						cnstr.Validate(bytes.NewBytes(given))
					})
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `Incorrect number value "not a number"`, func() {
			NewPrecision(bytes.NewBytes("2")).Validate(bytes.NewBytes("not a number"))
		})
	})
}

func TestPrecision_ASTNode(t *testing.T) {
	assert.Equal(t, schema.RuleASTNode{
		TokenType:  schema.TokenTypeNumber,
		Value:      "1",
		Properties: &schema.RuleASTNodes{},
		Source:     schema.RuleASTNodeSourceManual,
	}, NewPrecision(bytes.NewBytes("1")).ASTNode())
}
