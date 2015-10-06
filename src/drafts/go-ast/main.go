package ast

import (
	"go/parser"
	"go/token"
	"os"
)

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
