package jsoac

import (
	"fmt"
	"strings"

	schema "github.com/jsightapi/jsight-schema-core"
)

type Ref struct {
	userTypeName string
	nullable     bool
}

func newRef(astNode schema.ASTNode) Ref {
	return newRefFromUserTypeName(astNode.SchemaType, isNullable(astNode))
}

func newRefFromUserTypeName(name string, nullable bool) Ref {
	return Ref{
		userTypeName: name,
		nullable:     nullable,
	}
}

func (Ref) IsOpenAPINode() bool {
	return true
}

func (r Ref) MarshalJSON() ([]byte, error) {
	if r.nullable {
		return r.nullableJSON(), nil
	}

	return r.basicJSON(), nil
}

func (r Ref) basicJSON() []byte {
	s := fmt.Sprintf(`{"$ref": "#/components/schemas/%s"}`, strings.TrimLeft(r.userTypeName, "@"))
	return []byte(s)
}

func (r Ref) nullableJSON() []byte {
	s := fmt.Sprintf(`{"nullable": true, "allOf": [ %s ]}`, r.basicJSON())
	return []byte(s)
}
