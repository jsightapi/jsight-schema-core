package openapi

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
)

type Node struct {
	// key is the key for the object element.
	key string

	// type_ is an OpenAPI data type.
	type_ Type

	// example is an OpenAPI example.
	example Example

	// children represent available object properties or array items.
	children []Node
}

var _ json.Marshaler = Node{}
var _ json.Marshaler = &Node{}

func newNode(astNode schema.ASTNode) Node {
	switch astNode.TokenType {
	case schema.TokenTypeString:
		return newPrimitiveNode(TypeString, astNode.Value)
	case schema.TokenTypeBoolean:
		return newPrimitiveNode(TypeBoolean, astNode.Value)
	case schema.TokenTypeNumber:
		return newNumber(astNode)
	case schema.TokenTypeArray:
		return newArrayNode(astNode)
	case schema.TokenTypeObject:
		return newObjectNode(astNode)
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func newPrimitiveNode(t Type, ex string) Node {
	return Node{
		type_:   t,
		example: newExample(t, ex),
	}
}

func newNumber(astNode schema.ASTNode) Node {
	if astNode.SchemaType == "integer" {
		return newPrimitiveNode(TypeInteger, astNode.Value)
	}

	return newPrimitiveNode(TypeNumber, astNode.Value)
}

func newArrayNode(astNode schema.ASTNode) Node {
	o := Node{
		type_:    TypeArray,
		children: make([]Node, 0, len(astNode.Children)),
	}

	for _, an := range astNode.Children {
		o.children = append(o.children, newNode(an))
	}

	return o
}

func newObjectNode(astNode schema.ASTNode) Node {
	o := Node{
		type_:    TypeObject,
		children: make([]Node, 0, len(astNode.Children)),
	}

	for _, an := range astNode.Children {
		n := newNode(an)
		n.key = an.Key

		o.children = append(o.children, n)
	}

	return o
}

func (n Node) MarshalJSON() ([]byte, error) {
	switch n.type_ {
	case TypeObject:
		return n.marshalObjectNode()
	case TypeArray:
		return n.marshalArrayNode()
	default:
		return n.marshalPrimitiveNode()
	}
}

func (n Node) marshalPrimitiveNode() ([]byte, error) {
	var data struct {
		Type    string  `json:"type"`
		Example Example `json:"example"`
	}

	data.Type = n.type_.String()
	data.Example = n.example

	return json.Marshal(data)
}

func (n Node) marshalArrayNode() ([]byte, error) {
	var data struct {
		Type  string `json:"type"`
		Items []Node `json:"items"`
	}

	data.Type = n.type_.String()
	data.Items = n.children

	return json.Marshal(data)
}

func (n Node) marshalObjectNode() ([]byte, error) {
	var data struct {
		Type       string           `json:"type"`
		Properties ObjectProperties `json:"properties"`
	}

	data.Type = n.type_.String()
	data.Properties = ObjectProperties{
		properties: n.children,
	}

	return json.Marshal(data)
}
