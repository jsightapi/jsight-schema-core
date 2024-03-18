package info

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type ElementInfo struct {
	target ElementType
	node   *schema.ASTNode
}

func newElementInfo(t ElementType) ElementInfo {
	return ElementInfo{target: t}
}

func (e ElementInfo) Type() ElementType {
	return e.target
}

func (e ElementInfo) Children() []schema.ASTNode {
	return e.node.Children
}
func (e *ElementInfo) setASTNode(astNode schema.ASTNode) {
	e.node = &astNode
}
