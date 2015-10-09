package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func renderTemplate(in string) (string, error) {
	out, err := render(in, nil)
	if err != nil {
		return "", err
	}
	return out, nil
}

func render(t string, i interface{}) (string, error) {
	tpl, err := template.New(t).Funcs(funcs).Parse(t)
	if err != nil {
		log.Printf("error parsing")
		return "", err
	}
	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, i)
	return buf.String(), err
}

var funcs = template.FuncMap{
	"require": funcRequire,
	"code":    funcCode,
}

func funcRequire(in string) (string, error) {
	f, err := os.Open(in)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func funcCode(in string) string {
	lines := strings.Split(in, "\n")
	out := []string{}
	for _, l := range lines {
		out = append(out, "    "+l)
	}
	return strings.Join(out, "\n")
}
