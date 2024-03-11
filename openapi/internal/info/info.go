package info

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/openapi/internal/jsoac"
)

type Info struct {
	root *jschema.JSchema
	node schema.ASTNode
}

func NewInfo(s *jschema.JSchema, astNode schema.ASTNode) Info {
	return Info{
		root: s,
		node: astNode,
	}
}

func (i Info) astNode() schema.ASTNode {
	return i.node
}

func (i Info) SchemaObject() *jsoac.JSOAC {
	return jsoac.NewFromASTNode(i.node)
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
	return i.propertiesInfos(i.node)
}

func (i Info) propertiesInfos(astNode schema.ASTNode) *Properties {
	switch astNode.TokenType {
	case schema.TokenTypeObject:
		return i.propertiesInfosFromObject(astNode)
	case schema.TokenTypeShortcut:
		return i.propertiesInfosFromShortcut(astNode)
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func (i Info) propertiesInfosFromShortcut(astNode schema.ASTNode) *Properties {
	name := astNode.Value

	ut, ok := i.root.UserTypeCollection[name]
	if !ok {
		panic(errs.ErrUserTypeNotFound.F(name))
	}

	jut, ok := ut.(*jschema.JSchema)
	if !ok {
		panic(errs.ErrIncorrectUserType.F())
	}

	return i.propertiesInfos(jut.ASTNode)
}

func (i Info) propertiesInfosFromObject(astNode schema.ASTNode) *Properties {
	props := newProperties(len(astNode.Children))

	for _, child := range astNode.Children {
		props.append(child.Key, NewInfo(i.root, child))
	}

	return props
}
