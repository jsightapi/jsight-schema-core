package openapi

import (
	"encoding/json"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type OpenAPI struct {
	schema *jschema.JSchema // TODO it's necessary?
	root   Node
}

var _ json.Marshaler = OpenAPI{}
var _ json.Marshaler = &OpenAPI{}

func New(j *jschema.JSchema) OpenAPI {
	return OpenAPI{
		schema: j,
		root:   newNode(j.ASTNode),
	}
}

func (o OpenAPI) MarshalJSON() (b []byte, err error) {
	return o.root.MarshalJSON()
}
