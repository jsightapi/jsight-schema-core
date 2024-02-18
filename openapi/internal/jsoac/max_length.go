package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMaxLength(astNode schema.ASTNode, t OADType) *Example {
	if astNode.Rules.Has("maxLength") && t == OADTypeString {
		return exampleRef(newExample(astNode.Rules.GetValue("maxLength").Value, OADTypeInteger))
	}
	return nil
}
