package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/json"

	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestNewMin(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		ruleValue := bytes.NewBytes("3.14")
		cnstr := NewMin(ruleValue)

		expectedNumber, err := json.NewNumber(ruleValue)
		require.NoError(t, err)

		assert.Equal(t, ruleValue, cnstr.rawValue)
		assert.Equal(t, expectedNumber, cnstr.min)
		assert.False(t, cnstr.exclusive)
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `Incorrect number value "not a number"`, func() {
			NewMin(bytes.NewBytes("not a number"))
		})
	})
}

func TestMin_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(t, Min{}, json.TypeInteger, json.TypeFloat)
}

func TestMin_Type(t *testing.T) {
	assert.Equal(t, MinConstraintType, NewMin(bytes.NewBytes("1")).Type())
}

func TestMin_String(t *testing.T) {
	cc := map[bool]string{
		false: "min: 3.14",
		true:  "min: 3.14 (exclusive: true)",
	}

	for exclusive, expected := range cc {
		t.Run(expected, func(t *testing.T) {
			cnstr := NewMin(bytes.NewBytes("3.14"))
			cnstr.SetExclusive(exclusive)

			actual := cnstr.String()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMin_SetExclusive(t *testing.T) {
	cnstr := Min{}

	cnstr.SetExclusive(true)
	assert.True(t, cnstr.exclusive)

	cnstr.SetExclusive(false)
	assert.False(t, cnstr.exclusive)
}

func TestMin_Exclusive(t *testing.T) {
	cnstr := Min{}

	cnstr.exclusive = true
	assert.True(t, cnstr.Exclusive())

	cnstr.exclusive = false
	assert.False(t, cnstr.Exclusive())
}

func TestMin_Validate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		newMin := func(max string, exclusive bool) *Min {
			cnstr := NewMin(bytes.NewBytes(max))
			cnstr.SetExclusive(exclusive)
			return cnstr
		}

		cc := map[string]struct {
			cnstr *Min
			value string
			error string
		}{
			"3.14 >= 3.14": {
				cnstr: newMin("3.14", true),
				value: "3.14",
				error: "The value in the example violates the rule `\"min\": 3.14` (exclusive)",
			},
			"3.14 >= 2": {
				cnstr: newMin("3.14", true),
				value: "2",
				error: "The value in the example violates the rule `\"min\": 3.14` (exclusive)",
			},
			"3.14 > 3.14": {
				cnstr: newMin("3.14", false),
				value: "3.14",
			},
			"3.14 > 2": {
				cnstr: newMin("3.14", false),
				value: "2",
				error: "The value in the example violates the rule `\"min\": 3.14` ",
			},
			"3.14 >= 4": {
				cnstr: newMin("3.14", true),
				value: "4",
			},
			"3.14 > 4": {
				cnstr: newMin("3.14", false),
				value: "4",
			},
		}

		for name, c := range cc {
			t.Run(name, func(t *testing.T) {
				if c.error != "" {
					assert.PanicsWithError(t, c.error, func() {
						c.cnstr.Validate(bytes.NewBytes(c.value))
					})
				} else {
					assert.NotPanics(t, func() {
						c.cnstr.Validate(bytes.NewBytes(c.value))
					})
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `Incorrect number value "not a number"`, func() {
			NewMin(bytes.NewBytes("3")).Validate(bytes.NewBytes("not a number"))
		})
	})
}

func TestMin_ASTNode(t *testing.T) {
	assert.Equal(t, schema.RuleASTNode{
		TokenType:  schema.TokenTypeNumber,
		Value:      "1",
		Properties: &schema.RuleASTNodes{},
		Source:     schema.RuleASTNodeSourceManual,
	}, NewMin(bytes.NewBytes("1")).ASTNode())
}

func TestMin_Value(t *testing.T) {
	num, err := json.NewNumber(bytes.NewBytes("42"))
	require.NoError(t, err)

	cnstr := Min{
		min: num,
	}
	assert.Equal(t, num, cnstr.Value())
}
