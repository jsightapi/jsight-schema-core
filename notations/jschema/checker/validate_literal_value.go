package checker

import (
	"sort"

	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/errors"
	"github.com/jsightapi/jsight-schema-core/json"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema"
	"github.com/jsightapi/jsight-schema-core/notations/jschema/ischema/constraint"
)

func ValidateLiteralValue(node ischema.Node, jsonValue bytes.Bytes) {
	checkJsonType(node, jsonValue)

	// sorting to make it easier to debug the scheme if there are several errors in it
	m := node.ConstraintMap()
	l := m.Len()
	keys := make([]int, 0, l)

	m.EachSafe(func(k constraint.Type, _ constraint.Constraint) {
		keys = append(keys, int(k))
	})

	sort.Ints(keys)

	var isNullable bool
	if c, ok := m.Get(constraint.NullableConstraintType); ok {
		isNullable = c.(constraint.BoolKeeper).Bool()
	}

	for _, k := range keys {
		t := constraint.Type(k)
		c := m.GetValue(t)

		if _, ok := c.(*constraint.Enum); ok && isNullable && jsonValue.String() == "null" {
			// Handle cases like `null // {enum: [1, 2], nullable: true}`.
			continue
		}

		if v, ok := c.(constraint.LiteralValidator); ok {
			v.Validate(jsonValue)
		}
	}
}

func checkJsonType(node ischema.Node, value bytes.Bytes) {
	if node.Constraint(constraint.EnumConstraintType) != nil {
		return
	}

	jsonType := json.Guess(value).LiteralJsonType() // can panic
	schemaType := node.Type()
	if !(jsonType == schemaType ||
		(jsonType == json.TypeInteger && schemaType == json.TypeFloat) ||
		(jsonType == json.TypeNull && node.Constraint(constraint.NullableConstraintType) != nil)) {
		panic(errors.Format(errors.ErrInvalidValueType, jsonType.String(), schemaType.String()))
	}
}
