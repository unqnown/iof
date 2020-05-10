package iof

import (
	"os"

	encxml "encoding/xml"
)

// XML is Encoding implementation for "xml" data format.
const XML xml = ".xml"

var _ Encoding = XML

type xml string

func (xml) String() string { return string(XML) }

func (xml) Read(name string, v interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encxml.NewDecoder(f).Decode(v)
}

func (xml) Write(name string, v interface{}) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encxml.NewEncoder(f).Encode(v)
}

func (xml) Insert(name string, v interface{}) error {
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
	return encxml.NewEncoder(f).Encode(v)
}

func (xml) Update(name string, v interface{}) error {
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
	return encxml.NewEncoder(f).Encode(v)
}

func (x xml) Upsert(name string, v interface{}) error { return x.Write(name, v) }
