package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var REVISION string

func main() {
	var i interface{}

	if err := json.Unmarshal([]byte(REVISION), &i); err != nil {
		log.Fatalf("unable to unmarshal %q: %s", REVISION, err)
	}

	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", b)
}
