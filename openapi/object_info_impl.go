package openapi

import (
	"github.com/jsightapi/jsight-schema-core/openapi/info"
)

type ObjectInfoImpl struct {
	ElementInfo
}

func newObjectInfoImpl(ei ElementInfo) ObjectInfoImpl {
	return ObjectInfoImpl{ElementInfo: ei}
}

func (o ObjectInfoImpl) PropertiesInfos() []PropertyInfo {
	props := o.ElementInfo.(info.ElementInfo).Children()
	result := make([]PropertyInfo, len(props))

	for i, child := range props {
		result[i] = newPropertyInfoImpl(child)
	}

	return result
}
