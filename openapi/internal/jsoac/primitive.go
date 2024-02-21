package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	jstType schema.TokenType
	OADType OADType `json:"type"`
	Example Example `json:"example"`
	// optional fields
	Pattern          *Pattern  `json:"pattern,omitempty"`
	Format           *string   `json:"format,omitempty"`
	Enum             *Enum     `json:"enum,omitempty"`
	Minimum          *Number   `json:"minimum,omitempty"`
	Maximum          *Number   `json:"maximum,omitempty"`
	ExclusiveMinimum *bool     `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum *bool     `json:"exclusiveMaximum,omitempty"`
	MinLength        *int64    `json:"minLength,omitempty"`
	MaxLength        *int64    `json:"maxLength,omitempty"`
	MultipleOf       *float64  `json:"multipleOf,omitempty"`
	Nullable         *Nullable `json:"nullable,omitempty"`
}

func newPrimitive(t OADType, astNode schema.ASTNode) Primitive {
	if astNode.SchemaType == "integer" {
		t = OADTypeInteger
	}
	var p = Primitive{
		OADType:          t,
		Example:          newExample(astNode.Value, t),
		Pattern:          newPattern(astNode),
		Format:           newFormat(astNode),
		Enum:             newEnum(astNode, t),
		Minimum:          newMinimum(astNode, t),
		Maximum:          newMaximum(astNode, t),
		ExclusiveMinimum: newExclusiveMinimum(astNode, t),
		ExclusiveMaximum: newExclusiveMaximum(astNode, t),
		MinLength:        newMinLength(astNode, t),
		MaxLength:        newMaxLength(astNode, t),
		MultipleOf:       newMultipleOf(astNode, t),
		Nullable:         newNullable(astNode),
	}
	return p
}

func (p Primitive) JSightTokenType() schema.TokenType {
	return p.jstType
}
