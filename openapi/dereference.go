package openapi

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/openapi/info"
)

func Dereference(s schema.Schema) []ElementInfo {
	dd := info.Dereference(s)
	result := make([]ElementInfo, len(dd))

	for i, d := range dd {
		result[i] = d
	}

	return result
}
