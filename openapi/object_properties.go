package openapi

import (
	"encoding/json"

	"github.com/jsightapi/jsight-schema-core/internal/sync"
)

type ObjectProperties struct {
	properties []Node
}

var _ json.Marshaler = ObjectProperties{}
var _ json.Marshaler = &ObjectProperties{}

var objectPropertiesBufferPool = sync.NewBufferPool(512)

func (p ObjectProperties) MarshalJSON() ([]byte, error) {
	b := objectPropertiesBufferPool.Get()
	defer objectPropertiesBufferPool.Put(b)

	b.WriteByte('{')
	length := len(p.properties)
	for i, property := range p.properties {
		b.WriteByte('"')
		b.WriteString(property.key)
		b.WriteString(`":`)

		value, err := property.MarshalJSON()
		if err != nil {
			return nil, err
		}
		b.Write(value)

		if i+1 != length {
			b.WriteByte(',')
		}
	}
	b.WriteByte('}')
	return b.Bytes(), nil
}
