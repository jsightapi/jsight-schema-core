package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newEnumConst(astNode schema.ASTNode, t OADType) *Enum {
	if astNode.Rules.Has("const") && astNode.Rules.GetValue("const").Value == stringTrue {
		ex := newExample(astNode.Value, t)

		enum := makeEmptyEnum()
		enum.append(ex.value)
		return enum
	}
	return nil
}
