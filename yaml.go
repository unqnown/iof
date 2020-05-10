package iof

import (
	"os"

	encyaml "github.com/go-yaml/yaml"
)

// YAML is Encoding implementation for "yaml" data format.
const YAML yaml = ".yaml"

var _ Encoding = YAML

type yaml string

func (yaml) String() string { return string(YAML) }

func (yaml) Read(name string, v interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encyaml.NewDecoder(f).Decode(v)
}

func (yaml) Write(name string, v interface{}) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return encyaml.NewEncoder(f).Encode(v)
}

func (yaml) Insert(name string, v interface{}) error {
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
	return encyaml.NewEncoder(f).Encode(v)
}

func (yaml) Update(name string, v interface{}) error {
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
	return encyaml.NewEncoder(f).Encode(v)
}

func (y yaml) Upsert(name string, v interface{}) error { return y.Write(name, v) }
