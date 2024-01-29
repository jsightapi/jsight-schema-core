package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	jstType schema.TokenType
	OADType OADType `json:"type"`
	Example Example `json:"example"`
}

func newString(astNode schema.ASTNode) Primitive {
	return Primitive{
		OADType: OADTypeString,
		Example: newStringExample(astNode.Value),
	}
}

func newBoolean(astNode schema.ASTNode) Primitive {
	return newBasicNode(OADTypeBoolean, astNode.Value)
}

func newNumber(astNode schema.ASTNode) Primitive {
	if astNode.SchemaType == "integer" {
		return newBasicNode(OADTypeInteger, astNode.Value)
	}

	return newBasicNode(OADTypeNumber, astNode.Value)
}

func newBasicNode(t OADType, astValue string) Primitive {
	return Primitive{
		OADType: t,
		Example: newBasicExample(astValue),
	}
}

func (p Primitive) JSightTokenType() schema.TokenType {
	return p.jstType
}
