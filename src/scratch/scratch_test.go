package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestScratch(t *testing.T) {
	os.RemoveAll("tmp")
	defer os.RemoveAll("tmp")

	s := &Scratcher{Root: "tmp"}
	msg := &Message{}
	msg.Time = time.Date(2001, 2, 3, 4, 5, 0, 0, time.UTC)
	msg.Message = "first message"

	if err := s.Write(msg); err != nil {
		t.Fatal(err)
	}

	msg.Message = "second message"
	if err := s.Write(msg); err != nil {
		t.Fatal(err)
	}

	lines, err := readLines("tmp/2001/02/03/2001-02-03T04.log")
	if err != nil {
		t.Fatal(err)
	}
	if v, ex := len(lines), 2; v != ex {
		t.Errorf("expected to have %d lines, was %d", ex, v)
	}

	if v, ex := lines[0], "first message"; !strings.Contains(v, ex) {
		t.Errorf("expected first line %q to contain %q", v, ex)
	}

	if v, ex := lines[1], "second message"; !strings.Contains(v, ex) {
		t.Errorf("expected second line %q to contain %q", v, ex)
	}

}

func readLines(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n"), nil
}
