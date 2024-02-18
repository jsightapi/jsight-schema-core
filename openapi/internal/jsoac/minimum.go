package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMinimum(astNode schema.ASTNode, t OADType) *Example {
	if astNode.Rules.Has("min") && (t == OADTypeInteger || t == OADTypeNumber) {
		return exampleRef(newExample(astNode.Rules.GetValue("min").Value, t))
	}
	return nil
}
