package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMinLength(astNode schema.ASTNode, t OADType) *Example {
	if astNode.Rules.Has("minLength") && t == OADTypeString {
		return exampleRef(newExample(astNode.Rules.GetValue("minLength").Value, OADTypeInteger))
	}
	return nil
}
