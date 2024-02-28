package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

func newConst(astNode schema.ASTNode) *Enum {
	if astNode.Rules.Has("const") && astNode.Rules.GetValue("const").Value == stringTrue {
		ex := newExample(astNode.Value, isString(astNode))

		enum := makeEmptyEnum()
		enum.append(ex.value)
		return enum
	}
	return nil
}
