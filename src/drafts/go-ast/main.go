package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fs := token.NewFileSet()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	res, err := parser.ParseDir(fs, wd, nil, parser.ImportsOnly)
	if err != nil {
		return err
	}
	_ = res
	return nil
}
