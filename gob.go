package iof

import (
	"os"

	encgob "encoding/gob"
)

// GOB is Encoding implementation for "gob" data format.
const GOB gob = ".gob"

var _ Encoding = GOB

type gob string

func (gob) String() string { return string(GOB) }

func (gob) Read(name string, v interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encgob.NewDecoder(f).Decode(v)
}

func (gob) Write(name string, v interface{}) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encgob.NewEncoder(f).Encode(v)
}

func (gob) Insert(name string, v interface{}) error {
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
	return encgob.NewEncoder(f).Encode(v)
}

func (gob) Update(name string, v interface{}) error {
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
	return encgob.NewEncoder(f).Encode(v)
}

func (y gob) Upsert(name string, v interface{}) error { return y.Write(name, v) }
