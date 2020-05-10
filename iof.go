package iof

import (
	"errors"
	"fmt"
	"path/filepath"
)

var (
	// ErrAlreadyExists means that that resource of given name already exists.
	ErrAlreadyExists = errors.New("iof: already exists")
	// ErrNotExists means that that resource of given name not exists.
	ErrNotExists = errors.New("iof: not exists")
	// ErrNotFile means that resource of given name is not a file.
	ErrNotFile = errors.New("iof: not file")
)

// Reader is the interface that decodes content of a file
// to the given value.
type Reader interface {
	Read(name string, v interface{}) error
}

// Writer is the interface that creates a new file
// with the content of encoded value, regardless
// to file existence. If file already exists -
// it will be overrided with a new content.
type Writer interface {
	Write(name string, v interface{}) error
}

type ReadWriter interface {
	Reader
	Writer
}

// Encoding represents interface for file i/o workflow.
type Encoding interface {
	fmt.Stringer
	Reader
	Writer

	// Insert creates a new file with the content of encoded value.
	// If file already exists returns ErrAlreadyExists.
	Insert(name string, v interface{}) error
	// Update updates existing file with content of encoded value.
	// If file not exists returns ErrNotExists.
	Update(name string, v interface{}) error
	// Upsert is an alias for Write method.
	Upsert(name string, v interface{}) error
}

var encodings = map[string]Encoding{
	CSV.String():  CSV,
	GOB.String():  GOB,
	JSON.String(): JSON,
	XML.String():  XML,
	YAML.String(): YAML,
	".yml":        YAML,
}

// RegisterEncoding makes new Encoding available for
// resolving by file extension.
func RegisterEncoding(ext string, enc Encoding) { encodings[ext] = enc }

var defenc Encoding = JSON

// Read decodes content of a file to the given value
// via Encoding resolved from file's extension.
func Read(name string, v interface{}) error {
	enc, found := encodings[filepath.Ext(name)]
	if !found {
		enc = defenc
	}
	return enc.Read(name, v)
}

// Write encodes given value to a file via Encoding
// resolved from file's extension.
func Write(name string, v interface{}) error {
	enc, found := encodings[filepath.Ext(name)]
	if !found {
		enc = defenc
	}
	return enc.Write(name, v)
}

// Insert encodes given value to a new file via Encoding
// resolved from file's extension.
func Insert(name string, v interface{}) error {
	enc, found := encodings[filepath.Ext(name)]
	if !found {
		enc = defenc
	}
	return enc.Insert(name, v)
}

// Update encodes given value to the existing file via Encoding
// resolved from file's extension.
func Update(name string, v interface{}) error {
	enc, found := encodings[filepath.Ext(name)]
	if !found {
		enc = defenc
	}
	return enc.Update(name, v)
}

// Upsert encodes given value to a file via Encoding
// resolved from file's extension.
func Upsert(name string, v interface{}) error {
	enc, found := encodings[filepath.Ext(name)]
	if !found {
		enc = defenc
	}
	return enc.Upsert(name, v)
}
