package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/lexeme"

	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/fs"
)

func TestNewConstraintFromRule(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]struct {
			val          string
			expectedType Constraint
		}{
			"minLength":            {"1", &MinLength{}},
			"maxLength":            {"1", &MaxLength{}},
			"min":                  {"1", &Min{}},
			"max":                  {"1", &Max{}},
			"exclusiveMinimum":     {"true", &ExclusiveMinimum{}},
			"exclusiveMaximum":     {"true", &ExclusiveMaximum{}},
			"type":                 {"foo", &TypeConstraint{}},
			"precision":            {"1", &Precision{}},
			"optional":             {"true", &Optional{}},
			"minItems":             {"1", &MinItems{}},
			"maxItems":             {"1", &MaxItems{}},
			"additionalProperties": {"true", &AdditionalProperties{}},
			"nullable":             {"true", &Nullable{}},
			"regex":                {`"."`, &Regex{}},
			"const":                {"true", &Const{}},
		}

		for given, c := range cc {
			t.Run(given, func(t *testing.T) {
				constraint := NewConstraintFromRule(
					lexeme.NewLexEvent(
						lexeme.LiteralBegin,
						0,
						bytes.Index(len(given))-1,
						fs.NewFile("", given),
					),
					bytes.NewBytes(c.val),
					bytes.Bytes{},
				)

				assert.IsType(t, c.expectedType, constraint)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `ERROR (code 601): Unknown rule "invalid". See the list of all possible rules here: https://jsight.io/docs/jsight-schema-0-3#rules
	in line 1 on file 
	> invalid
	--^`, func() {
			const given = "invalid"

			NewConstraintFromRule(
				lexeme.NewLexEvent(
					lexeme.LiteralBegin,
					0,
					bytes.Index(len(given))-1,
					fs.NewFile("", given),
				),
				bytes.Bytes{},
				bytes.Bytes{},
			)
		})
	})
}
