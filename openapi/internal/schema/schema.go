package schema

import (
	"encoding/json"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type Schema struct {
	root Node
}

var _ json.Marshaler = Schema{}
var _ json.Marshaler = &Schema{}

func New(j *jschema.JSchema) Schema {
	return Schema{
		root: newNode(j.ASTNode),
	}
}

func (o Schema) MarshalJSON() (b []byte, err error) {
	return json.Marshal(o.root)
}
