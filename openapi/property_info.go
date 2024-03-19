package openapi

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/jsoac"
)

type PropertyInfo struct {
	node schema.ASTNode
}

var _ PropertyInformer = PropertyInfo{}
var _ PropertyInformer = (*PropertyInfo)(nil)

func newPropertyInfo(astNode schema.ASTNode) PropertyInfo {
	return PropertyInfo{node: astNode}
}

func (i PropertyInfo) Key() string {
	return i.node.Key
}

func (i PropertyInfo) Optional() bool {
	v, ok := i.node.Rules.Get("optional")
	if !ok {
		return false
	}
	return v.Value == "true"
}

func (i PropertyInfo) Annotation() string {
	return i.node.Comment
}

func (i PropertyInfo) SchemaObject() SchemaObject {
	return jsoac.NewFromASTNode(i.node)
}
