package openapi

type ObjectInfo struct {
	ElementInfo
}

var _ ObjectInformer = ObjectInfo{}
var _ ObjectInformer = (*ObjectInfo)(nil)

func newObjectInfo(t ElementType) ObjectInfo {
	return ObjectInfo{ElementInfo: newElementInfo(t)}
}

func (o ObjectInfo) PropertiesInfos() []PropertyInformer {
	props := o.ElementInfo.Children()
	result := make([]PropertyInformer, len(props))

	for i, child := range props {
		result[i] = newPropertyInfo(child)
	}

	return result
}
