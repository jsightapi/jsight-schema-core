package schema

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Object struct {
	Type_      Type             `json:"type"`
	Properties ObjectProperties `json:"properties"`
}

func newObject(astNode schema.ASTNode) Object {
	props := newObjectProperties(len(astNode.Children))

	for _, an := range astNode.Children {
		props.append(an.Key, newNode(an))
	}

	o := Object{
		Type_:      TypeObject,
		Properties: props,
	}

	return o
}

func (o Object) Type() Type {
	return o.Type_
}
