package openapi

import (
	"github.com/jsightapi/jsight-schema-core/openapi/info"
)

type ObjectInfoImpl struct {
	info.ElementInfo
}

func newObjectInfoImpl(ei info.ElementInfo) ObjectInfoImpl {
	return ObjectInfoImpl{ElementInfo: ei}
}

func (o ObjectInfoImpl) PropertiesInfos() []PropertyInfo {
	result := make([]PropertyInfo, len(o.Children()))

	for i, child := range o.Children() {
		result[i] = newPropertyInfoImpl(child)
	}

	return result
}
