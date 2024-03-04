package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Null struct {
	Example     *Example     `json:"example,omitempty"`
	Enum        *Enum        `json:"enum,omitempty"`
	Nullable    *Nullable    `json:"nullable,omitempty"`
	Description *Description `json:"description,omitempty"`
}

func newNull(astNode schema.ASTNode) Null {
	return Null{
		Example:     newExample(stringNull, false),
		Enum:        newEnum(astNode),
		Nullable:    newNullable(astNode),
		Description: newDescription(astNode),
	}
}

func (Null) IsOpenAPINode() bool {
	return true
}
