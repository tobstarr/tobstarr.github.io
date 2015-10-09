package main

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	in := `
	# Header 1

	{{ require "out.txt" | code }}

	# Header 2
	`

	out, err := renderTemplate(in)
	if err != nil {
		t.Fatal(err)
	}

	if v, ex := out, "    This is\n    out.txt"; !strings.Contains(v, ex) {
		t.Errorf("expected %q to contain %q", v, ex)
	}
}
