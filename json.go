package iof

import (
	"io"
	"os"

	encjson "encoding/json"
)

// JSON is Encoding implementation for "json" data format.
const JSON json = ".json"

var _ Encoding = JSON

type json string

func (json) String() string { return string(JSON) }

func (json) Read(name string, v interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encjson.NewDecoder(f).Decode(v)
}

func (j json) Write(name string, v interface{}) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return j.encode(f, v)
}

func (j json) Insert(name string, v interface{}) error {
	switch _, err := os.Stat(name); {
	case err == nil:
		return ErrAlreadyExists
	case os.IsNotExist(err):
		//
	default:
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	return j.encode(f, v)
}

func (j json) Update(name string, v interface{}) error {
	switch stat, err := os.Stat(name); {
	case err == nil:
		if stat.IsDir() {
			return ErrNotFile
		}
	case os.IsNotExist(err):
		return ErrNotExists
	default:
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	return j.encode(f, v)
}

func (j json) Upsert(name string, v interface{}) error { return j.Write(name, v) }

func (json) encode(w io.Writer, v interface{}) error {
	enc := encjson.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "	")
	return enc.Encode(v)
}
