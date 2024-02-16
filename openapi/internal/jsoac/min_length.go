package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"

	"strconv"
)

func newMinLength(astNode schema.ASTNode, t OADType) *int64 {
	if astNode.Rules.Has("minLength") && t == OADTypeString {
		v := astNode.Rules.GetValue("minLength").Value
		minLength, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		return int64Ref(minLength)
	}
	return nil
}
