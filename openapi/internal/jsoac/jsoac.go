package jsoac

// JSight schema to OpenAPi converter

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
)

type JSOAC struct {
	root Node
}

func New(j *jschema.JSchema) JSOAC {
	return JSOAC{
		root: newNode(j.ASTNode),
	}
}

func NewFromASTNode(a schema.ASTNode) JSOAC {
	return JSOAC{
		root: newNode(a),
	}
}

func (o JSOAC) JSON() (b []byte, err error) {
	return json.Marshal(o.root)
}
