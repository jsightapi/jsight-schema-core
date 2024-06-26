package constraint

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/json"

	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestNewConst(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]bool{
			"true":  true,
			"false": false,
		}

		for given, expected := range cc {
			t.Run(given, func(t *testing.T) {
				c := fakeConst(given, "foo")
				assert.Equal(t, expected, c.apply)
				assert.Equal(t, "foo", c.nodeValue.String())
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, `Invalid value in the "const" rule. Learn about the rules here: https://jsight.io/docs/jsight-schema-0-3#rules`, func() {
			fakeConst("foo", "")
		})
	})
}

func TestConst_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(
		t,
		Const{},
		json.TypeUndefined,
		json.TypeString,
		json.TypeInteger,
		json.TypeFloat,
		json.TypeBoolean,
		json.TypeNull,
		json.TypeMixed,
	)
}

func TestConst_Type(t *testing.T) {
	assert.Equal(t, ConstConstraintType, Const{}.Type())
}

func TestConst_String(t *testing.T) {
	cc := map[string]string{
		"false": "const: false",
		"true":  "const: true",
	}

	for given, expected := range cc {
		t.Run(given, func(t *testing.T) {
			assert.Equal(t, expected, fakeConst(given, "").String())
		})
	}
}

func TestConst_Bool(t *testing.T) {
	cc := map[string]bool{
		"false": false,
		"true":  true,
	}

	for given, expected := range cc {
		t.Run(given, func(t *testing.T) {
			assert.Equal(t, expected, fakeConst(given, "").Bool())
		})
	}
}

func TestConst_Validate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("apply - true", func(t *testing.T) {
			fakeConst("true", "foo").Validate(bytes.NewBytes("foo"))
		})

		t.Run("apply - false", func(t *testing.T) {
			t.Run("valid", func(t *testing.T) {
				fakeConst("false", "foo").Validate(bytes.NewBytes("foo"))
			})

			t.Run("invalid", func(t *testing.T) {
				fakeConst("false", "foo").Validate(bytes.NewBytes("bar"))
			})
		})
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, "Does not match expected value (foo)", func() {
			fakeConst("true", "foo").Validate(bytes.NewBytes("bar"))
		})
	})
}

func TestConst_ASTNode(t *testing.T) {
	cc := []bool{true, false}

	for _, c := range cc {
		t.Run(strconv.FormatBool(c), func(t *testing.T) {
			assert.Equal(t, schema.RuleASTNode{
				TokenType:  schema.TokenTypeBoolean,
				Value:      strconv.FormatBool(c),
				Properties: &schema.RuleASTNodes{},
				Source:     schema.RuleASTNodeSourceManual,
			}, Const{apply: c}.ASTNode())
		})
	}
}

func fakeConst(v, nv string) *Const {
	return NewConst(bytes.NewBytes(v), bytes.NewBytes(nv))
}
