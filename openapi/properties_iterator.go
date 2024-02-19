package openapi

type PropertiesIterator interface {
	Rewind()
	Next() bool
	GetKey() string
	GetInfo() SchemaInfo
}
