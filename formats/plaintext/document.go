package plaintext

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/jsightapi/jsight-schema-core/lexeme"
)

type Document struct {
	file *fs.File
}

var _ schema.Document = &Document{}

// New creates a JSON document with specified name and content.
func New[T bytes.ByteKeeper](name string, content T) *Document {
	return &Document{
		file: fs.NewFile(name, content),
	}
}

func (d *Document) Len() (uint, error) {
	return uint(d.file.Content().Len()), nil
}

func (d *Document) Content() bytes.Bytes {
	return d.file.Content()
}

// NextLexeme doesn't make sense for this structure
func (d *Document) NextLexeme() (lexeme.LexEvent, error) {
	return lexeme.NewLexEvent(
		lexeme.LiteralEnd,
		0,
		bytes.Index(d.file.Content().Len()-1),
		d.file,
	), nil
}

// Check doesn't make sense for this structure
func (d *Document) Check() error {
	return nil
}
