package openapi

import "github.com/jsightapi/jsight-schema-core/openapi/internal/info"

type SchemaInfoImpl struct {
	info.Info
}

var _ SchemaInfo = SchemaInfoImpl{}
var _ SchemaInfo = (*SchemaInfoImpl)(nil)

func newSchemaInfoImpl(i info.Info) SchemaInfoImpl {
	return SchemaInfoImpl{i}
}

func (i SchemaInfoImpl) SchemaObject() SchemaObject {
	return i.Info.SchemaObject()
}

func (i SchemaInfoImpl) PropertiesInfos() PropertiesIterator {
	return newPropertiesIteratorImpl(i.Info.PropertiesInfos())
}
