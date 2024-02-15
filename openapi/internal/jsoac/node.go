package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
)

type Node interface {
	JSightTokenType() schema.TokenType
}

func newNode(astNode schema.ASTNode) Node {
	switch astNode.TokenType {
	case schema.TokenTypeNumber:
		return newPrimitive(OADTypeNumber, astNode)
	case schema.TokenTypeString:
		return newPrimitive(OADTypeString, astNode)
	case schema.TokenTypeBoolean:
		return newPrimitive(OADTypeBoolean, astNode)
	case schema.TokenTypeArray:
		return newArray(astNode)
	case schema.TokenTypeObject:
		return newObject(astNode)
	case schema.TokenTypeNull:
		return newNull()
	// TODO case schema.TokenTypeShortcut:
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}
