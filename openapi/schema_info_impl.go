package openapi

import "github.com/jsightapi/jsight-schema-core/openapi/internal/info"

type SchemaInfoImpl struct {
	info.Info
}

func (i SchemaInfoImpl) SchemaObject() SchemaObject {
	return i.Info.SchemaObject()
}

func (i SchemaInfoImpl) PropertiesInfos() PropertiesIterator {
	return i.Info.PropertiesInfos()
}
