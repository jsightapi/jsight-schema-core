package openapi

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/jsoac"
)

type PropertyInfoImpl struct {
	node schema.ASTNode
}

func newPropertyInfoImpl(astNode schema.ASTNode) PropertyInfoImpl {
	return PropertyInfoImpl{node: astNode}
}

func (i PropertyInfoImpl) Key() string {
	return i.node.Key
}

func (i PropertyInfoImpl) Optional() bool {
	v, ok := i.node.Rules.Get("optional")
	if !ok {
		return false
	}
	return v.Value == "true"
}

func (i PropertyInfoImpl) Annotation() string {
	return i.node.Comment
}

func (i PropertyInfoImpl) SchemaObject() SchemaObject {
	return jsoac.NewFromASTNode(i.node)
}
