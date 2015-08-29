package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	if err := run(l); err != nil {
		log.Fatal(err)
	}
}

type Job struct {
	Out   string
	Parts []string
}

func (j *Job) Exec(l Logger) error {
	return writeFile(l, j.Out, j.Parts...)
}

func run(l *log.Logger) error {
	jobs := []*Job{
		{"index.html", []string{"tpl/header.tpl", "tpl/index.tpl", "tpl/footer.tpl"}},
		{"dotfiles/index.html", []string{"tpl/header.tpl", "tpl/dotfiles/index.tpl", "tpl/footer.tpl"}},
		{"dotfiles/vimrc.conf", []string{"dotfiles/vimrc"}},
	}
	for _, j := range jobs {
		if err := j.Exec(l); err != nil {
			return err
		}
	}
	return nil
}

type Logger interface {
	Printf(string, ...interface{})
}

func writeFile(l Logger, out string, files ...string) error {
	start := time.Now()
	buf := &bytes.Buffer{}
	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		_, err = io.Copy(buf, f)
		if err != nil {
			return err
		}
		f.Close()
	}
	tmp := out + ".tmp." + strconv.FormatInt(time.Now().UTC().UnixNano(), 36)
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	_, err = io.Copy(f, buf)
	if err != nil {
		return err
	}
	err = os.Rename(tmp, out)
	if err != nil {
		return err
	}
	l.Printf("write file %s in %.6f", out, time.Since(start).Seconds())
	return nil
}
