package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMaximum(astNode schema.ASTNode, t OADType) *Number {
	if astNode.Rules.Has("max") && (t == OADTypeInteger || t == OADTypeNumber) {
		return numberRef(newNumber(astNode.Rules.GetValue("max").Value))
	}
	return nil
}
