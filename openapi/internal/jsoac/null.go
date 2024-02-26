package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Null struct {
	jstType  schema.TokenType
	Example  Example   `json:"example"`
	Enum     *Enum     `json:"enum"`
	Nullable *Nullable `json:"nullable,omitempty"`
}

func newNull(astNode schema.ASTNode) Null {
	return Null{
		jstType:  schema.TokenTypeNull,
		Example:  newExample(stringNull, OADTypeInteger),
		Enum:     newEnum(astNode, OADTypeInteger),
		Nullable: newNullable(astNode),
	}
}

func (n Null) JSightTokenType() schema.TokenType {
	return n.jstType
}
