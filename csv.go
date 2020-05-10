package iof

import (
	"os"

	"github.com/gocarina/gocsv"
)

// CSV is Encoding implementation for "csv" data format.
const CSV csv = ".csv"

var _ Encoding = CSV

type csv string

func (csv) String() string { return string(CSV) }

func (csv) Read(name string, v interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return gocsv.Unmarshal(f, v)
}

func (csv) Write(name string, v interface{}) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return gocsv.Marshal(v, f)
}

func (csv) Insert(name string, v interface{}) error {
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
	return gocsv.Marshal(v, f)
}

func (csv) Update(name string, v interface{}) error {
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
	return gocsv.Marshal(v, f)
}

func (y csv) Upsert(name string, v interface{}) error { return y.Write(name, v) }
