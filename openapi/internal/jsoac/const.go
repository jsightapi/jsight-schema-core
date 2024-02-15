package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newConst(astNode schema.ASTNode, p *Primitive) {
	if astNode.Rules.GetValue("const").Value == "true" {
		// TODO replace with Enum
		var value = newBasicExample(p.OADType, astNode.Value)
		var enum [1]Example
		enum[0] = value
		p.CEnum = &enum
	}
}
