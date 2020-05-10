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
