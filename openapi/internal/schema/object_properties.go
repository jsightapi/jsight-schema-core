package schema

import (
	"encoding/json"
	"github.com/jsightapi/jsight-schema-core/internal/sync"
)

type ObjectProperties struct {
	properties []Property
}

type Property struct {
	key   string
	value Node
}

var _ json.Marshaler = ObjectProperties{}
var _ json.Marshaler = &ObjectProperties{}

var objectPropertiesBufferPool = sync.NewBufferPool(512)

func newObjectProperties(len int) ObjectProperties {
	return ObjectProperties{properties: make([]Property, 0, len)}
}

func (op *ObjectProperties) append(key string, value Node) {
	p := Property{
		key:   key,
		value: value,
	}
	op.properties = append(op.properties, p)
}

func (op ObjectProperties) MarshalJSON() ([]byte, error) {
	b := objectPropertiesBufferPool.Get()
	defer objectPropertiesBufferPool.Put(b)

	b.WriteByte('{')
	length := len(op.properties)
	for i, property := range op.properties {
		b.WriteByte('"')
		b.WriteString(property.key)
		b.WriteString(`":`)

		value, err := json.Marshal(property.value)
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
