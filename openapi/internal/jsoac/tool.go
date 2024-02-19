package jsoac

import (
	"github.com/jsightapi/jsight-schema-core/internal/sync"

	"github.com/jsightapi/jsight-schema-core/notations/jschema"

	"strconv"

	"github.com/stretchr/testify/require"

	"testing"
)

const stringTrue = "true"

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

func numberRef(n Number) *Number {
	return &n
}

func int64Ref(i int64) *int64 {
	return &i
}

func int64RefByString(s string) *int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	return int64Ref(value)
}
