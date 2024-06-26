package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/json"

	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestNewMax(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		ruleValue := bytes.NewBytes("3.14")
		cnstr := NewMax(ruleValue)

		expectedNumber, err := json.NewNumber(ruleValue)
		require.NoError(t, err)

		assert.Equal(t, ruleValue, cnstr.rawValue)
		assert.Equal(t, expectedNumber, cnstr.max)
		assert.False(t, cnstr.exclusive)
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `Incorrect number value "not a number"`, func() {
			NewMax(bytes.NewBytes("not a number"))
		})
	})
}

func TestMax_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(t, Max{}, json.TypeInteger, json.TypeFloat)
}

func TestMax_Type(t *testing.T) {
	assert.Equal(t, MaxConstraintType, NewMax(bytes.NewBytes("1")).Type())
}

func TestMax_String(t *testing.T) {
	cc := map[bool]string{
		false: "max: 3.14",
		true:  "max: 3.14 (exclusive: true)",
	}

	for exclusive, expected := range cc {
		t.Run(expected, func(t *testing.T) {
			cnstr := NewMax(bytes.NewBytes("3.14"))
			cnstr.SetExclusive(exclusive)

			actual := cnstr.String()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMax_SetExclusive(t *testing.T) {
	cnstr := Max{}

	cnstr.SetExclusive(true)
	assert.True(t, cnstr.exclusive)

	cnstr.SetExclusive(false)
	assert.False(t, cnstr.exclusive)
}

func TestMax_Exclusive(t *testing.T) {
	cnstr := Max{}

	cnstr.exclusive = true
	assert.True(t, cnstr.Exclusive())

	cnstr.exclusive = false
	assert.False(t, cnstr.Exclusive())
}

func TestMax_Validate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		newMax := func(max string, exclusive bool) *Max {
			cnstr := NewMax(bytes.NewBytes(max))
			cnstr.SetExclusive(exclusive)
			return cnstr
		}

		cc := map[string]struct {
			cnstr *Max
			value string
			error string
		}{
			"3.14 <= 3.14": {
				cnstr: newMax("3.14", true),
				value: "3.14",
				error: "The value in the example violates the rule `\"max\": 3.14` (exclusive)",
			},
			"3.14 <= 2": {
				cnstr: newMax("3.14", true),
				value: "2",
			},
			"3.14 < 3.14": {
				cnstr: newMax("3.14", false),
				value: "3.14",
			},
			"3.14 < 2": {
				cnstr: newMax("3.14", false),
				value: "2",
			},
			"3.14 <= 4": {
				cnstr: newMax("3.14", true),
				value: "4",
				error: "The value in the example violates the rule `\"max\": 3.14` (exclusive)",
			},
			"3.14 < 4": {
				cnstr: newMax("3.14", false),
				value: "4",
				error: "The value in the example violates the rule `\"max\": 3.14` ",
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
			NewMax(bytes.NewBytes("3")).Validate(bytes.NewBytes("not a number"))
		})
	})
}

func TestMax_ASTNode(t *testing.T) {
	assert.Equal(t, schema.RuleASTNode{
		TokenType:  schema.TokenTypeNumber,
		Value:      "1",
		Properties: &schema.RuleASTNodes{},
		Source:     schema.RuleASTNodeSourceManual,
	}, NewMax(bytes.NewBytes("1")).ASTNode())
}

func TestMax_Value(t *testing.T) {
	num, err := json.NewNumber(bytes.NewBytes("42"))
	require.NoError(t, err)

	cnstr := Max{
		max: num,
	}
	assert.Equal(t, num, cnstr.Value())
}
