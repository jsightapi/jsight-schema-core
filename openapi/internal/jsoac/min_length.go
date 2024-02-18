package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMinLength(astNode schema.ASTNode, t OADType) *int64 {
	if astNode.Rules.Has("minLength") && t == OADTypeString {
		return int64RefByString(astNode.Rules.GetValue("minLength").Value)
	}
	return nil
}
