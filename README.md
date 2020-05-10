### iof [![GoDoc](https://godoc.org/github.com/unqnown/iof?status.svg)](https://godoc.org/github.com/unqnown/iof)

iof simplifies i/o operations with files of different data-serialization formats.

### features

- designed to facilitate the process of reading/writing files with simple API:

```go
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
```

- decreases amount of mechanical code like:

```go
f, err := os.Create("file.json")
if err != nil {
    return err
}
defer f.Close()
enc := json.NewEncoder(f)
enc.SetEscapeHTML(false)
enc.SetIndent("", "	")
return enc.Encode(f)
```

- allows performing i/o operations according to file's extension.

There are few other methods which provide more accurately reading/writing functionality:
- Insert - writes content only to a new file; returns error if file already exists;
- Update - writes content only to existing file; returns error if file not exists;
- Upsert - does same as `Write` method.

### installation

Standard `go get`:

```shell script
go get github.com/unqnown/iof
````

### usage & example

```go
package iof_test

import (
	"fmt"
	"log"

	"github.com/unqnown/iof"
)

type Data struct {
	V interface{}
}

func Example() {
	if err := iof.JSON.Write("test.json", Data{V: "ʕ•ᴥ•ʔ"}); err != nil {
		log.Fatal(err)
	}
	var read Data
	if err := iof.JSON.Read("test.json", &read); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", read.V)
	// output: ʕ•ᴥ•ʔ
}

func Example_Global() {
	if err := iof.Write("test.yaml", Data{V: "ʕ•ᴥ•ʔ"}); err != nil {
		log.Fatal(err)
	}
	var read Data
	if err := iof.Read("test.yaml", &read); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", read.V)
	// output: ʕ•ᴥ•ʔ
}
```
