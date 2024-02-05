package openapi

import "github.com/jsightapi/jsight-schema-core/openapi/internal/info"

type SchemaInfoImpl struct {
	info info.Info
}

func newSchemaInfoImpl(i info.Info) SchemaInfoImpl {
	return SchemaInfoImpl{info: i}
}

func (i SchemaInfoImpl) SchemaObject() SchemaObject {
	return i.info.SchemaObject()
}

func (i SchemaInfoImpl) Optional() bool {
	return i.info.Optional()
}

func (i SchemaInfoImpl) Annotation() string {
	return i.info.Annotation()
}

func (i SchemaInfoImpl) NestedObjectProperties() []SchemaInfo {
	a := make([]SchemaInfo, 0, len(i.info.NestedObjectProperties()))

	for _, child := range i.info.NestedObjectProperties() {
		a = append(a, newSchemaInfoImpl(child))
	}

	return a
}
