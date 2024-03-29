package jsoac

import (
	"encoding/json"

	"github.com/jsightapi/jsight-schema-core/openapi/internal"

	schema "github.com/jsightapi/jsight-schema-core"

	"github.com/jsightapi/jsight-schema-core/errs"
)

type additionalPropertiesMode int

const (
	additionalPropertiesNull additionalPropertiesMode = iota
	additionalPropertiesFalse
	additionalPropertiesArray
	additionalPropertiesPrimitive
	additionalPropertiesFormat
	additionalPropertiesUserType
	additionalPropertiesAnyOf
	additionalPropertiesObject
)

type AdditionalProperties struct {
	mode         additionalPropertiesMode
	oadType      *OADType
	format       string
	userTypeName string
	node         schema.ASTNode
}

var _ json.Marshaler = &AdditionalProperties{}

func newAdditionalProperties(astNode schema.ASTNode) *AdditionalProperties {
	isKeyShortcutWithAP := isKeyShortcutWithAP(astNode)
	haveStringAdditionalProperties := astNode.Rules.Has(stringAdditionalProperties)
	var r *schema.RuleASTNode = nil
	var addProp *AdditionalProperties = nil
	if haveStringAdditionalProperties {
		r = refRuleASTNode(astNode.Rules.GetValue(stringAdditionalProperties))
		switch r.TokenType {
		case schema.TokenTypeBoolean:
			addProp = newBooleanAdditionalProperties(r)
			if !isKeyShortcutWithAP {
				return addProp
			}
		case schema.TokenTypeString:
			addProp = newStringAdditionalProperties(r)
			if !isKeyShortcutWithAP {
				return addProp
			}
		default:
			panic(errs.ErrRuntimeFailure.F())
		}
	}

	// key shortcut type
	for _, an := range astNode.Children {
		if an.IsKeyShortcut {
			if haveStringAdditionalProperties {
				if (r.TokenType == schema.TokenTypeBoolean && r.Value == stringTrue) ||
					(r.TokenType == schema.TokenTypeString && r.Value == stringAny) {
					return nil
				}
			}
			return newAnyOfAdditionalProperties(astNode)
		}
	}
	if addProp != nil {
		return addProp
	}

	// The additionalProperties JSight rule is missing
	return newFalseAdditionalProperties()
}

// check is astNode have childrens with some key as shortcut and with additional properties
func isKeyShortcutWithAP(astNode schema.ASTNode) bool {
	for _, an := range astNode.Children {
		if an.IsKeyShortcut && astNode.Rules.Has(stringAdditionalProperties) {
			return true
		}
	}
	return false
}

func newAnyOfAdditionalProperties(node schema.ASTNode) *AdditionalProperties {
	oadType := OADTypeObject
	return &AdditionalProperties{
		mode:    additionalPropertiesAnyOf,
		oadType: &oadType,
		node:    node,
	}
}

func newStringAdditionalProperties(r *schema.RuleASTNode) *AdditionalProperties {
	if r.Value == stringNull {
		return &AdditionalProperties{mode: additionalPropertiesNull}
	}

	if r.Value == stringArray {
		return &AdditionalProperties{mode: additionalPropertiesArray}
	}

	if r.Value == stringObject {
		return &AdditionalProperties{mode: additionalPropertiesObject}
	}

	if r.Value == stringAny {
		return nil
	}

	if r.Value[0] == '@' {
		return &AdditionalProperties{mode: additionalPropertiesUserType, userTypeName: r.Value}
	}

	t := oadTypeFromSchemaType(r.Value)
	f := internal.FormatFromSchemaType(r.Value)

	if f == nil {
		return &AdditionalProperties{
			mode:    additionalPropertiesPrimitive,
			oadType: &t,
		}
	}

	return &AdditionalProperties{
		mode:    additionalPropertiesFormat,
		oadType: &t,
		format:  *f,
	}
}

func newBooleanAdditionalProperties(r *schema.RuleASTNode) *AdditionalProperties {
	if r.Value == stringFalse {
		return newFalseAdditionalProperties()
	}
	return nil // JSight additionalProperties: true
}

func newFalseAdditionalProperties() *AdditionalProperties {
	return &AdditionalProperties{
		mode: additionalPropertiesFalse,
	}
}

func (a *AdditionalProperties) MarshalJSON() ([]byte, error) {
	switch a.mode {
	case additionalPropertiesFalse:
		return a.booleanJSON()
	case additionalPropertiesNull:
		return a.nullJSON()
	case additionalPropertiesArray:
		return a.arrayJSON()
	case additionalPropertiesObject:
		return a.objectJSON()
	case additionalPropertiesFormat:
		return a.formatJSON()
	case additionalPropertiesPrimitive:
		return a.primitiveJSON()
	case additionalPropertiesUserType:
		return a.userTypeJSON()
	case additionalPropertiesAnyOf:
		return a.anyOfJSON(a.node)
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func (a *AdditionalProperties) anyOfJSON(node schema.ASTNode) ([]byte, error) {
	var items []any

	if node.Rules.Has(stringAdditionalProperties) {
		r := node.Rules.GetValue(stringAdditionalProperties)
		if r.TokenType == schema.TokenTypeString {
			additionalAnyJSONObject := makeAdditionalAnyJSONObjects(r)
			items = append(items, additionalAnyJSONObject)
		}
	}

	for _, astNode := range node.Children {
		if astNode.Key != "" && astNode.Key[0] == '@' {
			node := newNode(astNode)
			items = append(items, node)
		}
	}

	data := struct {
		Items []any `json:"anyOf"`
	}{
		Items: items,
	}
	m, err := json.Marshal(data)
	return m, err
}

func (a *AdditionalProperties) arrayJSON() ([]byte, error) {
	data := struct {
		OADType OADType        `json:"type"`
		Items   map[string]any `json:"items"`
	}{
		OADType: OADTypeArray,
		Items:   map[string]any{},
	}
	return json.Marshal(data)
}

func (a *AdditionalProperties) objectJSON() ([]byte, error) {
	data := struct {
		OADType              OADType        `json:"type"`
		Properties           map[string]any `json:"properties"`
		AdditionalProperties bool           `json:"additionalProperties"`
	}{
		OADType:              OADTypeObject,
		Properties:           map[string]any{},
		AdditionalProperties: false,
	}
	return json.Marshal(data)
}

func (a *AdditionalProperties) booleanJSON() ([]byte, error) {
	return []byte(stringFalse), nil
}

func (a *AdditionalProperties) nullJSON() ([]byte, error) {
	return []byte(`{ "enum": [null] }`), nil
}

func (a *AdditionalProperties) primitiveJSON() ([]byte, error) {
	data := struct {
		OADType OADType `json:"type"`
	}{
		OADType: *a.oadType,
	}
	return json.Marshal(data)
}

func (a *AdditionalProperties) formatJSON() ([]byte, error) {
	data := struct {
		OADType OADType `json:"type"`
		Format  string  `json:"format"`
	}{
		OADType: *a.oadType,
		Format:  a.format,
	}
	return json.Marshal(data)
}

func (a *AdditionalProperties) userTypeJSON() ([]byte, error) {
	ref := newRefFromUserTypeName(a.userTypeName, false)
	return ref.MarshalJSON()
}
