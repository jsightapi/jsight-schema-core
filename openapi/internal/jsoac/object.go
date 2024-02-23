package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Object struct {
	cap                  int
	jstType              schema.TokenType
	OADType              OADType          `json:"type"`
	Properties           ObjectProperties `json:"properties"`
	AdditionalProperties bool             `json:"additionalProperties"`
	Required             []string         `json:"required,omitempty"`
	AllOf                *AllOf           `json:"allOf,omitempty"`
}

func newObject(astNode schema.ASTNode) Object {
	o := Object{
		cap:                  len(astNode.Children),
		OADType:              OADTypeObject,
		Properties:           newObjectProperties(len(astNode.Children)),
		AdditionalProperties: false, // TODO rule additionalProperties
		Required:             nil,
		AllOf:                newAllOf(astNode),
	}

	for _, an := range astNode.Children {
		o.appendProperty(an)
	}

	return o
}

func (o *Object) appendProperty(astNode schema.ASTNode) {
	key := astNode.Key
	value := newNode(astNode)

	o.Properties.append(key, value)

	if !astNode.Rules.Has("optional") || astNode.Rules.GetValue("optional").Value == stringFalse {
		o.appendToRequired(key)
	}
}

func (o *Object) appendToRequired(key string) {
	o.initRequiredIfNecessary()
	o.Required = append(o.Required, key)
}

func (o *Object) initRequiredIfNecessary() {
	if o.Required == nil {
		o.Required = make([]string, 0, o.cap)
	}
}

func (o Object) JSightTokenType() schema.TokenType {
	return o.jstType
}
