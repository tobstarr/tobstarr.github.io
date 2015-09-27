package main

import (
	"go/parser"
	"go/token"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fs := token.NewFileSet()
	res, err := parser.ParseDir(fs, d, nil, parser.ImportsOnly)
	if err != nil {
		return err
	}
	return nil
}
