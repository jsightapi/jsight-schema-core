package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	Type_   Type    `json:"type"`
	Example Example `json:"example"`
}

func newString(astNode schema.ASTNode) Primitive {
	return Primitive{
		Type_:   TypeString,
		Example: newStringExample(astNode.Value),
	}
}

func newBoolean(astNode schema.ASTNode) Primitive {
	return newBasicNode(TypeBoolean, astNode.Value)
}

func newNumber(astNode schema.ASTNode) Primitive {
	if astNode.SchemaType == "integer" {
		return newBasicNode(TypeInteger, astNode.Value)
	}

	return newBasicNode(TypeNumber, astNode.Value)
}

func newBasicNode(t Type, astValue string) Primitive {
	return Primitive{
		Type_:   t,
		Example: newBasicExample(astValue),
	}
}

func (p Primitive) Type() Type {
	return p.Type_
}
