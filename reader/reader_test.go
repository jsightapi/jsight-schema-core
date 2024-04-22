package reader

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/jsightapi/jsight-schema-core/test"
)

func TestRead(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		filename := filepath.Join(test.GetProjectRoot(), "testdata", "boolean.jschema")
		expected := bytes.NewBytes(`true // Schema containing a literal example`)

		file := fs.NewFile(filename, expected)

		assert.Equal(t, file, Read(filename))
	})

	t.Run("negative", func(t *testing.T) {
		assert.PanicsWithError(t, "open not_existing_file.jst: The system cannot find the file specified.", func() {
			Read("not_existing_file.jst")
		})
	})
}
