package jsoac

import (
	"bytes"

	"fmt"

	schema "github.com/jsightapi/jsight-schema-core"

	"github.com/jsightapi/jsight-schema-core/errs"

	"strings"
)

type SimpleByteArray []byte

type AdditionalPropertiesAnyJsonItem struct {
	Type                 *string            `json:"type,omitempty"`
	Items                *SimpleByteArray   `json:"items,omitempty"`
	Format               *string            `json:"format,omitempty"`
	Ref                  *string            `json:"$ref,omitempty"`
	Properties           *SimpleByteArray   `json:"properties,omitempty"`
	AdditionalProperties *bool              `json:"additionalProperties,omitempty"`
	Enum                 *[]SimpleByteArray `json:"enum,omitempty"`
	Example              *SimpleByteArray   `json:"example,omitempty"`
}

func makeAdditionalAnyJSONObjects(r schema.RuleASTNode) AdditionalPropertiesAnyJsonItem {
	var s AdditionalPropertiesAnyJsonItem
	var emptyByteObject SimpleByteArray = []byte("{}")

	switch r.Value {
	case stringString, stringInteger, stringBoolean:
		s = AdditionalPropertiesAnyJsonItem{
			Type: stringRef(r.Value),
		}
	case stringFloat:
		s = AdditionalPropertiesAnyJsonItem{
			Type: stringRef(stringNumber),
		}
	case stringArray:
		s = AdditionalPropertiesAnyJsonItem{
			Type:  stringRef(r.Value),
			Items: &emptyByteObject,
		}
	case stringObject:
		s = AdditionalPropertiesAnyJsonItem{
			Type:                 stringRef(stringObject),
			Properties:           &emptyByteObject,
			AdditionalProperties: boolRef(false),
		}
	case stringNull:
		var nullBytes SimpleByteArray = []byte(stringNull)
		simpleEnum := []SimpleByteArray{
			nullBytes,
		}
		s = AdditionalPropertiesAnyJsonItem{
			Enum:    &simpleEnum,
			Example: &nullBytes,
		}
	case stringDate:
		s = AdditionalPropertiesAnyJsonItem{
			Type:   stringRef(stringString),
			Format: stringRef(stringDate),
		}
	case stringDatetime:
		s = AdditionalPropertiesAnyJsonItem{
			Type:   stringRef(stringString),
			Format: stringRef("date-time"),
		}
	case stringEmail:
		s = AdditionalPropertiesAnyJsonItem{
			Type:   stringRef(stringString),
			Format: stringRef(stringEmail),
		}
	case stringUri:
		s = AdditionalPropertiesAnyJsonItem{
			Type:   stringRef(stringString),
			Format: stringRef(stringUri),
		}
	case stringUuid:
		s = AdditionalPropertiesAnyJsonItem{
			Type:   stringRef(stringString),
			Format: stringRef(stringUuid),
		}
	default:
		if r.Value[0] == '@' {
			s = AdditionalPropertiesAnyJsonItem{
				Ref: stringRef(fmt.Sprintf(`#/components/schemas/%s`, strings.TrimLeft(r.Value, "@"))),
			}
		} else {
			panic(errs.ErrRuntimeFailure.F()) // FIXME: may be: s = AdditionalPropertiesAnyJsonItem{}
		}
	}
	return s
}

func (s SimpleByteArray) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(s) == 0 {
		buf.Write([]byte(stringNull))
	} else {
		buf.Write(s)
	}
	return buf.Bytes(), nil
}
