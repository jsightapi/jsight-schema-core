package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newExclusiveMaximum(astNode schema.ASTNode, t OADType) *bool {
	if astNode.Rules.Has("exclusiveMaximum") && (t == OADTypeInteger || t == OADTypeNumber) {
		b := astNode.Rules.GetValue("exclusiveMaximum").Value == stringTrue
		if b {
			return &b
		}
		return nil
	}
	return nil
}