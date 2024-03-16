package info

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/openapi"
)

type ReferenceTargetInfo struct {
	target  openapi.ReferenceTargetType
	astNode *schema.ASTNode
}

func newReferenceTargetInfo(t openapi.ReferenceTargetType) ReferenceTargetInfo {
	return ReferenceTargetInfo{target: t}
}

func (r *ReferenceTargetInfo) SetASTNode(astNode schema.ASTNode) {
	r.astNode = &astNode
}

func (r ReferenceTargetInfo) ReferenceTargetType() openapi.ReferenceTargetType {
	return r.target
}
