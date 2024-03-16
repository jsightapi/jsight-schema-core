package info

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/notations/regex"
	"github.com/jsightapi/jsight-schema-core/openapi"
)

type dereference struct {
	src  schema.Schema
	root *jschema.JSchema
	lst  ReferenceTargetInfoList
}

func Dereference(s schema.Schema) []ReferenceTargetInfo {
	d := dereference{
		src:  s,
		root: nil,
		lst:  newReferenceTargetInfoList(),
	}

	return d.run()
}

func (d *dereference) run() []ReferenceTargetInfo {
	switch t := d.src.(type) {
	case *jschema.JSchema:
		d.root = t
		return d.dereferenceJSchema(t.ASTNode)
	case *regex.RSchema:
		return d.dereferenceRSchema()
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func (d *dereference) result() []ReferenceTargetInfo {
	return d.lst.list()
}

func (d *dereference) appendInfo(t openapi.ReferenceTargetType) *ReferenceTargetInfo {
	info := newReferenceTargetInfo(t)
	return d.lst.append(info)
}

func (d *dereference) dereferenceRSchema() []ReferenceTargetInfo {
	_ = d.appendInfo(openapi.ReferenceTargetTypeRegex)
	return d.result()
}

func (d *dereference) dereferenceJSchema(astNode schema.ASTNode) []ReferenceTargetInfo {
	// TODO OR
	switch astNode.TokenType {
	case schema.TokenTypeObject:
		info := d.appendInfo(openapi.ReferenceTargetTypeObject)
		info.SetASTNode(astNode)
		return d.result()
	case schema.TokenTypeShortcut:
		d.processShortcut(astNode)
		return d.result()
	default:
		panic(errs.ErrRuntimeFailure.F())
	}
}

func (d *dereference) processShortcut(astNode schema.ASTNode) {
	name := astNode.Value

	ut, ok := d.root.UserTypeCollection[name]
	if !ok {
		panic(errs.ErrUserTypeNotFound.F(name))
	}

	elements := Dereference(ut)

	for _, r := range elements {
		d.lst.append(r)
	}
}
