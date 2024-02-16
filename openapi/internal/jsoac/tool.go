package jsoac

import (
	"github.com/jsightapi/jsight-schema-core/internal/sync"
	"github.com/jsightapi/jsight-schema-core/notations/jschema"

	"github.com/stretchr/testify/require"

	"testing"
)

var bufferPool = sync.NewBufferPool(1024)

func quotedBytes(s string) []byte {
	bb := make([]byte, 0, len(s)+2)

	bb = append(bb, '"')
	bb = append(bb, []byte(s)...)
	bb = append(bb, '"')

	return bb
}

func jsightToOpenAPI(t *testing.T, jsight, openapi string) {
	j := jschema.New("TestSchemaName", jsight)
	err := j.Check()
	require.NoError(t, err)

	o := New(j)
	json, err := o.MarshalJSON()
	require.NoError(t, err)

	require.JSONEq(t, openapi, string(json), "Actual: "+string(json))
}

func strRef(s string) *string {
	return &s
}

func exampleRef(ex Example) *Example {
	return &ex
}
