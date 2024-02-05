package info

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/jsoac"
)

type Info struct {
	jschema *jschema.JSchema
	node    *schema.ASTNode // nullable
}

func New(s *jschema.JSchema) Info {
	return Info{jschema: s}
}

func newFromASTNode(s *jschema.JSchema, a *schema.ASTNode) Info {
	return Info{jschema: s, node: a}
}

func (i Info) astNode() *schema.ASTNode {
	if i.node != nil {
		return i.node
	}
	return &i.jschema.ASTNode
}

func (i Info) SchemaObject() *jsoac.JSOAC {
	return jsoac.New(i.jschema)
}

func (i Info) Optional() bool {
	v, ok := i.astNode().Rules.Get("optional")
	if !ok {
		return false
	}
	return v.Value == "true"
}

func (i Info) Annotation() string {
	return i.astNode().Comment
}

func (i Info) NestedObjectProperties() []Info {
	node := i.astNode()
	pp := make([]Info, 0, len(node.Children))

	if node.TokenType == schema.TokenTypeObject {
		for _, child := range node.Children {
			pp = append(pp, newFromASTNode(i.jschema, &child)) // #nosec G601 (nolint)
		}
	}

	return pp
}
