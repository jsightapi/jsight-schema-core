package kit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/test"

	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/fs"
)

func TestJSchemaError_preparation(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("not prepared", func(t *testing.T) {
			e := JSchemaError{file: fs.NewFile("", "123456")}
			e.preparation()

			assert.EqualValues(t, 6, e.length)
			assert.EqualValues(t, '\n', e.nl)
		})
	})

	t.Run("negative", func(t *testing.T) {
		test.PanicsWithErr(t, errs.ErrRuntimeFailure.F(), func() {
			(&JSchemaError{}).preparation()
		})
	})
}

type testValidResult struct {
	index  bytes.Index
	begin  bytes.Index
	end    bytes.Index
	str    string
	line   uint
	column uint
}
type testData struct {
	source string
	valid  []testValidResult
}

var data = []testData{
	{
		"ABC",
		[]testValidResult{
			{0, 0, 3, "ABC", 1, 1}, // A
			{1, 0, 3, "ABC", 1, 2}, // B
			{2, 0, 3, "ABC", 1, 3}, // C
		},
	},
	{
		"AB\n\nCD\n",
		[]testValidResult{
			{0, 0, 2, "AB", 1, 1}, // A
			{1, 0, 2, "AB", 1, 2}, // B
			{2, 0, 2, "AB", 1, 3}, // \n
			{3, 3, 3, "", 2, 1},   // \n
			{4, 4, 6, "CD", 3, 1}, // C
			{5, 4, 6, "CD", 3, 2}, // D
			{6, 4, 6, "CD", 3, 3}, // \n
		},
	},
	{
		"AB\r\rCD\r",
		[]testValidResult{
			{0, 0, 2, "AB", 1, 1}, // A
			{1, 0, 2, "AB", 1, 2}, // B
			{2, 0, 2, "AB", 1, 3}, // \r
			{3, 3, 3, "", 2, 1},   // \r
			{4, 4, 6, "CD", 3, 1}, // C
			{5, 4, 6, "CD", 3, 2}, // D
			{6, 4, 6, "CD", 3, 3}, // \r
		},
	},
	{
		"AB\r\n\r\nCD\r\n",
		[]testValidResult{
			{0, 0, 2, "AB", 1, 1}, // A
			{1, 0, 2, "AB", 1, 2}, // B
			{2, 0, 2, "AB", 1, 3}, // \r
			{3, 0, 2, "AB", 1, 4}, // \n
			{4, 4, 4, "", 2, 1},   // \r
			{5, 4, 4, "", 2, 2},   // \n
			{6, 6, 8, "CD", 3, 1}, // C
			{7, 6, 8, "CD", 3, 2}, // D
			{8, 6, 8, "CD", 3, 3}, // \r
			{9, 6, 8, "CD", 3, 4}, // \n
		},
	},
	{
		"AB\n\r\n\rCD\n\r",
		[]testValidResult{
			{0, 0, 2, "AB", 1, 1}, // A
			{1, 0, 2, "AB", 1, 2}, // B
			{2, 0, 2, "AB", 1, 3}, // \r
			{3, 0, 2, "AB", 1, 4}, // \n
			{4, 4, 4, "", 2, 1},   // \r
			{5, 4, 4, "", 2, 2},   // \n
			{6, 6, 8, "CD", 3, 1}, // C
			{7, 6, 8, "CD", 3, 2}, // D
			{8, 6, 8, "CD", 3, 3}, // \r
			{9, 6, 8, "CD", 3, 4}, // \n
		},
	},
	{
		"\n\n\n",
		[]testValidResult{
			{0, 0, 0, "", 1, 1},
			{1, 1, 1, "", 2, 1},
			{2, 2, 2, "", 3, 1},
		},
	},
	{
		"\nA\nB\n",
		[]testValidResult{
			{0, 0, 0, "", 1, 1},
			{1, 1, 2, "A", 2, 1},
			{2, 1, 2, "A", 2, 2},
			{3, 3, 4, "B", 3, 1},
			{4, 3, 4, "B", 3, 2},
		},
	},
	{
		"",
		[]testValidResult{
			{0, 0, 0, "", 0, 0},
			{1, 0, 0, "", 0, 0},
			{2, 0, 0, "", 0, 0},
		},
	},
}

func TestJSchemaError_lineBeginning(t *testing.T) {
	for _, d := range data {
		for _, v := range d.valid {
			t.Run(fmt.Sprintf("%s %d", d.source, v.index), func(t *testing.T) {
				file := fs.NewFile("", d.source)

				e := newFakeJSchemaError(file, v.index)

				begin := e.lineBeginning()
				assert.Equal(t, v.begin, begin)
			})
		}
	}
}

func TestJSchemaError_lineEnd(t *testing.T) {
	for _, d := range data {
		for _, v := range d.valid {
			t.Run(fmt.Sprintf("%s %d", d.source, v.index), func(t *testing.T) {
				file := fs.NewFile("", d.source)

				e := newFakeJSchemaError(file, v.index)

				end := e.lineEnd()
				assert.Equal(t, v.end, end)
			})
		}
	}
}

func TestNewJSchemaError_Line(t *testing.T) {
	for _, d := range data {
		for _, v := range d.valid {
			t.Run(fmt.Sprintf("%s %d", d.source, v.index), func(t *testing.T) {
				file := fs.NewFile("", d.source)

				e := newFakeJSchemaError(file, v.index)

				n := e.Line()
				assert.Equal(t, v.line, n)
			})
		}
	}
}

func TestNewJSchemaError_Column(t *testing.T) {
	for _, d := range data {
		for _, v := range d.valid {
			t.Run(fmt.Sprintf("%s %d", d.source, v.index), func(t *testing.T) {
				file := fs.NewFile("", d.source)

				e := newFakeJSchemaError(file, v.index)

				n := e.Column()
				assert.Equal(t, v.column, n)
			})
		}
	}
}

func TestJSchemaError_SourceSubString(t *testing.T) {
	for _, d := range data {
		for _, v := range d.valid {
			t.Run(fmt.Sprintf("%s %d", d.source, v.index), func(t *testing.T) {
				file := fs.NewFile("", d.source)

				e := newFakeJSchemaError(file, v.index)

				str := e.SourceSubString()
				assert.Equal(t, v.str, str)
			})
		}
	}

	t.Run("too long source substring", func(t *testing.T) {
		file := fs.NewFile("", strings.Repeat("123456789 ", 100))

		e := newFakeJSchemaError(file, 0)

		assert.Len(t, e.SourceSubString(), 200)
	})
}

func newFakeJSchemaError(f *fs.File, idx bytes.Index) JSchemaError {
	e := JSchemaError{}
	e.SetFile(f)
	e.SetIndex(idx)
	e.preparation()
	return e
}
