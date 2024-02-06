package openapi

import "github.com/jsightapi/jsight-schema-core/openapi/internal/info"

type PropertiesIterator interface {
	Rewind()
	Next() bool
	GetKey() string
	GetInfo() info.Info
}

var _ PropertiesIterator = (*info.Properties)(nil)
