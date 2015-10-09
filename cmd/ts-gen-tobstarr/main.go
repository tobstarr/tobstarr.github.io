package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tobstarr/tobstarr.github.io/cmd/ts-gen-tobstarr/Godeps/_workspace/src/github.com/shurcooL/github_flavored_markdown"
)

var sources = map[string][]Source{
	"docker.html":                          Layout(MarkdownSource("src/articles/docker/index.md")),
	"dotfiles.html":                        Layout(FileSource("src/dotfiles/index.tpl")),
	"dotfiles/vimrc.conf":                  FileSources("dotfiles/vimrc"),
	"gnupg.html":                           Layout(MarkdownSource("src/articles/gnupg/index.md")),
	"id_rsa.pub":                           FileSources("src/id_rsa.pub"),
	"index.html":                           Layout(FileSource("src/index.tpl")),
	"cheats.html":                          Layout(FileSource("src/cheats.tpl")),
	"tobstarr.gpg":                         FileSources("src/tobstarr.gpg"),
	"versions.html":                        Layout(MarkdownSource("src/versions.md")),
	"python-web-server.html":               Layout(MarkdownSource("src/python-web-server/index.md")),
	"geminabox.sh":                         FileSources("src/geminabox/geminabox.sh"),
	"speed-up-bundler-with-geminabox.html": Layout(MarkdownSource("src/speed-up-bundler-with-geminabox/index.md")),
}

func main() {
	l := log.New(os.Stderr, "", 0)
	if err := run(l); err != nil {
		log.Fatal(err)
	}
}

func FileSources(paths ...string) (list []Source) {
	for _, p := range paths {
		list = append(list, FileSource(p))
	}
	return list

}

func loadHTMLFiles() (list map[string]struct{}, err error) {
	list = map[string]struct{}{}
	return list, filepath.Walk(".", func(p string, info os.FileInfo, err error) error {
		switch {
		case info.IsDir():
			switch p {
			case ".git", "cmd":
				return filepath.SkipDir
			}
			log.Printf("checking %s", p)
			return nil
		case strings.HasSuffix(p, ".html"):
			list[p] = struct{}{}
		}
		return nil
	})
}

func Layout(s Source) []Source {
	return []Source{FileSource("src/header.tpl"), s, FileSource("src/footer.tpl")}
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
	start := time.Now()
	m, err := loadHTMLFiles()
	if err != nil {
		return err
	}
	for path, sources := range sources {
		if err := writeFile(l, path, sources...); err != nil {
			return err
		}
		delete(m, path)
	}
	for k := range m {
		l.Printf("deleting file %s", k)
		os.RemoveAll(k)
	}
	l.Printf("finished in %.06f", time.Since(start).Seconds())
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
