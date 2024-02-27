package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Any struct {
	Example     *Example     `json:"example,omitempty"`
	Nullable    *Nullable    `json:"nullable,omitempty"`
	Description *Description `json:"description,omitempty"`
}

func newAny(astNode schema.ASTNode) Any {
	a := Any{
		Nullable:    newNullable(astNode),
		Description: newDescription(astNode),
	}

	switch astNode.TokenType {
	case schema.TokenTypeString:
		ex := newExample(astNode.Value, OADTypeString)
		a.Example = &ex
	case schema.TokenTypeNumber, schema.TokenTypeBoolean, schema.TokenTypeNull:
		ex := newExample(astNode.Value, OADTypeNumber)
		a.Example = &ex
	}

	return a
}

func (Any) IsOpenAPINode() bool {
	return true
}
