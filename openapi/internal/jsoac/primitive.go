package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Primitive struct {
	OADType          *OADType     `json:"type,omitempty"`
	Example          Example      `json:"example"`
	Pattern          *Pattern     `json:"pattern,omitempty"`
	Format           *string      `json:"format,omitempty"`
	Enum             *Enum        `json:"enum,omitempty"`
	Minimum          *Number      `json:"minimum,omitempty"`
	Maximum          *Number      `json:"maximum,omitempty"`
	ExclusiveMinimum *bool        `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum *bool        `json:"exclusiveMaximum,omitempty"`
	MinLength        *int64       `json:"minLength,omitempty"`
	MaxLength        *int64       `json:"maxLength,omitempty"`
	MultipleOf       *float64     `json:"multipleOf,omitempty"`
	Nullable         *Nullable    `json:"nullable,omitempty"`
	Description      *Description `json:"description,omitempty"`
}

func oadType(schemaType string, t OADType) *OADType {
	if schemaType == stringEnum {
		return nil
	}
	return &t
}

func newPrimitive(astNode schema.ASTNode) Primitive {
	t := newOADType(astNode)
	var p = Primitive{
		OADType:          oadType(astNode.SchemaType, t),
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
		Description:      newDescription(astNode),
	}
	return p
}

func (Primitive) IsOpenAPINode() bool {
	return true
}
