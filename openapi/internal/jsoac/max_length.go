package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMaxLength(astNode schema.ASTNode, t OADType) *int64 {
	if astNode.Rules.Has("maxLength") && t == OADTypeString {
		return int64RefByString(astNode.Rules.GetValue("maxLength").Value)
	}
	return nil
}
