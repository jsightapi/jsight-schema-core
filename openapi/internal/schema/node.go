package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
)

type Node interface {
	Type() Type
}

func newNode(astNode schema.ASTNode) Node {
	switch astNode.TokenType {
	case schema.TokenTypeString:
		return newString(astNode)
	case schema.TokenTypeBoolean:
		return newBoolean(astNode)
	case schema.TokenTypeNumber:
		return newNumber(astNode)
	case schema.TokenTypeArray:
		return newArray(astNode)
	case schema.TokenTypeObject:
		return newObject(astNode)
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}
