package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Object struct {
	jstType    schema.TokenType
	OADType    OADType          `json:"type"`
	Properties ObjectProperties `json:"properties"`
}

func newObject(astNode schema.ASTNode) Object {
	props := newObjectProperties(len(astNode.Children))

	for _, an := range astNode.Children {
		props.append(an.Key, newNode(an))
	}

	o := Object{
		OADType:    OADTypeObject,
		Properties: props,
	}

	return o
}

func (o Object) JSightTokenType() schema.TokenType {
	return o.jstType
}
