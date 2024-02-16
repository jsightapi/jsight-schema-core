package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMaximum(astNode schema.ASTNode, t OADType) *Example {
	if astNode.Rules.Has("max") && (t == OADTypeInteger || t == OADTypeNumber) {
		return exampleRef(newExample(astNode.Rules.GetValue("max").Value, t))
	}
	return nil
}
