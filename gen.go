package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shurcooL/github_flavored_markdown"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	if err := run(l); err != nil {
		log.Fatal(err)
	}
}

type Job struct {
	Out     string
	Sources []Source
}

func (j *Job) Exec(l Logger) error {
	return writeFile(l, j.Out, j.Sources...)
}

func FileSources(paths ...string) (list []Source) {
	for _, p := range paths {
		list = append(list, FileSource(p))
	}
	return list

}

var sources = map[string][]Source{
	"dotfiles.html":       Layout(FileSource("tpl/dotfiles/index.tpl")),
	"dotfiles/vimrc.conf": FileSources("dotfiles/vimrc"),
	"index.html":          Layout(FileSource("tpl/index.tpl")),
	"versions.html":       Layout(MarkdownSource("tpl/versions.md")),
}

func Layout(s Source) []Source {
	return []Source{FileSource("tpl/header.tpl"), s, FileSource("tpl/footer.tpl")}
}

func MarkdownSource(path string) Source {
	return func(w io.Writer) error {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		b = github_flavored_markdown.Markdown(b)
		_, err = w.Write(b)
		return err
	}
}

func run(l *log.Logger) error {
	github_flavored_markdown.Markdown([]byte{})
	for path, sources := range sources {
		if err := writeFile(l, path, sources...); err != nil {
			return err
		}
	}
	return nil
}

func FileSource(path string) Source {
	return func(w io.Writer) error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	}
}

type Source func(w io.Writer) error

type Logger interface {
	Printf(string, ...interface{})
}

func writeFile(l Logger, out string, sources ...Source) error {
	start := time.Now()
	buf := &bytes.Buffer{}
	for _, src := range sources {
		if err := src(buf); err != nil {
			return err
		}
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
