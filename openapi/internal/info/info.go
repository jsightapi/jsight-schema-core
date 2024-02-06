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

func NewInfo(s *jschema.JSchema) Info {
	return Info{jschema: s}
}

func newInfoFromASTNode(s *jschema.JSchema, a schema.ASTNode) Info {
	return Info{jschema: s, node: &a}
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

func (i Info) PropertiesInfos() *Properties {
	node := i.astNode()
	props := newProperties(len(node.Children))

	if node.TokenType == schema.TokenTypeObject {
		for _, child := range node.Children {
			props.append(child.Key, newInfoFromASTNode(i.jschema, child))
		}
	}

	return props
}
