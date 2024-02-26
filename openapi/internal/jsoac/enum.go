package jsoac

import (
	"encoding/json"

	schema "github.com/jsightapi/jsight-schema-core"
)

type Enum struct {
	list [][]byte
}

var _ json.Marshaler = Enum{}
var _ json.Marshaler = &Enum{}

func newEnum(astNode schema.ASTNode, t OADType) *Enum {
	if enum := newEnumConst(astNode, t); enum != nil {
		return enum
	}
	// there will be other enums
	if astNode.Rules.Has(stringEnum) &&
		astNode.Rules.GetValue(stringEnum).TokenType == stringArray &&
		0 < len(astNode.Rules.GetValue(stringEnum).Items) {
		enum := makeEmptyEnum()
		for _, s := range astNode.Rules.GetValue(stringEnum).Items {
			valueType := t
			if s.TokenType == "string" {
				valueType = OADTypeString
			}
			ex := newExample(s.Value, valueType)
			enum.append(ex.value)
		}
		return enum
	}
	if astNode.SchemaType == stringNull || (astNode.TokenType == stringNull && astNode.SchemaType == stringAny) {
		enum := makeEmptyEnum()
		ex := newExample(stringNull, OADTypeInteger)
		enum.append(ex.value)
		return enum
	}
	return nil
}

func makeEmptyEnum() *Enum {
	return &Enum{
		list: make([][]byte, 0, 3),
	}
}

func (e *Enum) append(b []byte) {
	e.list = append(e.list, b)
}

func (e Enum) MarshalJSON() ([]byte, error) {
	b := bufferPool.Get()
	defer bufferPool.Put(b)

	b.WriteByte('[')
	for i, item := range e.list {
		b.Write(item)

		if i+1 != len(e.list) {
			b.WriteByte(',')
		}
	}
	b.WriteByte(']')

	return b.Bytes(), nil
}
