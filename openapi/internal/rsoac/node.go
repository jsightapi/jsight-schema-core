package rsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
)

type Node interface {
}

func newNode(astNode schema.ASTNode) Node {
	return newRegexString(astNode)
}
