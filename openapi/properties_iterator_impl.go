package openapi

import "github.com/jsightapi/jsight-schema-core/openapi/internal/info"

type PropertiesIteratorImpl struct {
	*info.Properties
}

var _ PropertiesIterator = PropertiesIteratorImpl{}
var _ PropertiesIterator = (*PropertiesIteratorImpl)(nil)

func newPropertiesIteratorImpl(p *info.Properties) PropertiesIteratorImpl {
	return PropertiesIteratorImpl{p}
}

func (p PropertiesIteratorImpl) GetInfo() SchemaInfo {
	return newSchemaInfoImpl(p.Properties.GetInfo())
}
