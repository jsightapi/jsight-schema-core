package jsoac

// JSight schema to OpenAPi converter

import (
	"encoding/json"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type JSOAC struct {
	root        Node
	description string
}

func New(j *jschema.JSchema) *JSOAC {
	return &JSOAC{
		root: newNode(j.ASTNode),
	}
}

func (o *JSOAC) SetDescription(s string) {
	o.description = s
}

func (o JSOAC) MarshalJSON() (b []byte, err error) {
	return json.Marshal(o.root)
}
