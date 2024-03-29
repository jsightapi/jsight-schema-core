package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/test"

	"github.com/jsightapi/jsight-schema-core/json"

	"github.com/jsightapi/jsight-schema-core/bytes"
)

func TestUUID_IsJsonTypeCompatible(t *testing.T) {
	testIsJsonTypeCompatible(t, UUID{}, json.TypeString)
}

func TestUUID_Type(t *testing.T) {
	assert.Equal(t, UuidConstraintType, NewUuid().Type())
}

func TestUUID_String(t *testing.T) {
	assert.Equal(t, "uuid", NewUuid().String())
}

func TestUUID_Validate(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		var tests = []string{
			`550e8400-e29b-41d4-a716-446655440000`,
			`urn:uuid:550e8400-e29b-41d4-a716-446655440000`,
			`URN:UUID:550e8400-e29b-41d4-a716-446655440000`,
			`{550e8400-e29b-41d4-a716-446655440000}`,
			`550e8400e29b41d4a716446655440000`,
			`aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee`,
			`AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE`,
		}

		for _, value := range tests {
			t.Run(value, func(t *testing.T) {
				NewUuid().Validate(bytes.NewBytes(value))
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		var tests = map[string]string{
			"":      "Invalid UUID length: 0",
			"12":    "Invalid UUID length: 2",
			"1.2":   "Invalid UUID length: 3",
			"true":  "Invalid UUID length: 4",
			"false": "Invalid UUID length: 5",
			"null":  "Invalid UUID length: 4",
			`"ABC"`: "Invalid UUID length: 3",
			// leading symbol " "
			" 550e8400e29b41d4a716446655440000": "Invalid UUID length: 33",
			// leading symbol " "
			" 550e8400-e29b-41d4-a716-446655440000": "Invalid UUID length: 37",
			// trailing symbol " "
			"550e8400e29b41d4a716446655440000 ": "Invalid UUID length: 33",
			// trailing symbol " "
			"550e8400-e29b-41d4-a716-446655440000 ": "Invalid UUID length: 37",
			// leading  and trailing symbol " "
			" 550e8400e29b41d4a716446655440000 ": "Invalid UUID length: 34",
			// leading  and trailing symbol " "
			" 550e8400-e29b-41d4-a716-446655440000 ": "Invalid prefix: braces expected",
			// additional trailing symbol "9"
			"550e8400e29b41d4a7164466554400009": "Invalid UUID length: 33",
			// invalid symbol "-" location
			"550e840-0e29b-41d4-a716-446655440000": "Invalid UUID format",
			// invalid symbol "z"
			"z50e8400-e29b-41d4-a716-446655440000":          "Invalid UUID format",
			"not:uuid:550e8400-e29b-41d4-a716-446655440000": `Invalid URN prefix: "not:uuid:"`,
		}

		for given, expected := range tests {
			t.Run(given, func(t *testing.T) {
				test.PanicsWithErr(t, errs.ErrInvalidUUID.F(expected), func() {
					NewUuid().Validate(bytes.NewBytes(given))
				})
			})
		}
	})
}

func TestUUID_ASTNode(t *testing.T) {
	assert.Equal(t, newEmptyRuleASTNode(), UUID{}.ASTNode())
}
