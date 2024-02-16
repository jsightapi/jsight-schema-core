package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"

	"strconv"
)

func newMaxLength(astNode schema.ASTNode, t OADType) *int64 {
	if astNode.Rules.Has("maxLength") && t == OADTypeString {
		v := astNode.Rules.GetValue("maxLength").Value
		maxLength, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		return int64Ref(maxLength)
	}
	return nil
}
