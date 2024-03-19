package openapi

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/notations/regex"
	"github.com/jsightapi/jsight-schema-core/openapi/internal"
)

type dereference struct {
	userTypes map[string]schema.Schema
	result    *elementInfoList
}

func Dereference(s schema.Schema) []ElementInformer {
	d := dereference{
		userTypes: nil,
		result:    newElementInfoList(),
	}

	if st, ok := s.(*jschema.JSchema); ok {
		d.userTypes = st.UserTypeCollection
	}

	d.schema(s)

	return d.result.list()
}

func (d dereference) schema(s schema.Schema) {
	switch st := s.(type) {
	case *jschema.JSchema:
		d.jSchema(st.ASTNode)
	case *regex.RSchema:
		d.rSchema()
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func (d dereference) rSchema() {
	info := newElementInfo(ElementTypeRegex)
	d.result.append(info)
}

func (d dereference) jSchema(astNode schema.ASTNode) {
	if rule, ok := astNode.Rules.Get("or"); ok {
		for _, item := range rule.Items {
			d.orItem(item)
		}
		return
	}

	if astNode.SchemaType == "any" {
		info := newElementInfo(ElementTypeAny)
		info.setASTNode(astNode)
		d.result.append(info)
		return
	}

	switch astNode.TokenType {
	case schema.TokenTypeNumber, schema.TokenTypeString, schema.TokenTypeBoolean, schema.TokenTypeNull:
		info := newElementInfo(ElementTypeScalar)
		info.setASTNode(astNode)
		d.result.append(info)
	case schema.TokenTypeObject:
		info := newObjectInfo(ElementTypeObject)
		info.setASTNode(astNode)
		d.result.append(info)
	case schema.TokenTypeArray:
		info := newElementInfo(ElementTypeArray)
		info.setASTNode(astNode)
		d.result.append(info)
	case schema.TokenTypeShortcut:
		name := astNode.Value
		d.userType(name)
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func (d dereference) userType(name string) {
	ut, ok := d.userTypes[name]
	if !ok {
		panic(errs.ErrUserTypeNotFound.F(name))
	}

	d.schema(ut)
}

func (d dereference) orItem(r schema.RuleASTNode) {
	mockAstNode := internal.RuleToASTNode(r)
	d.jSchema(mockAstNode)
}
