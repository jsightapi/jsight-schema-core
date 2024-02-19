package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newMinimum(astNode schema.ASTNode, t OADType) *Number {
	if astNode.Rules.Has("min") && (t == OADTypeInteger || t == OADTypeNumber) {
		return numberRef(newNumber(astNode.Rules.GetValue("min").Value))
	}
	return nil
}
