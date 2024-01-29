package schema

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type Schema struct {
	root Node
}

var _ json.Marshaler = Schema{}
var _ json.Marshaler = &Schema{}

func NewFromJSchema(j *jschema.JSchema) Schema {
	return Schema{
		root: newNode(j.ASTNode),
	}
}

func NewFromASTNode(a schema.ASTNode) Schema {
	return Schema{
		root: newNode(a),
	}
}

func (o Schema) MarshalJSON() (b []byte, err error) {
	return json.Marshal(o.root)
}
