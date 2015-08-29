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

func run(l *log.Logger) error {
	return writeFile("index.html", "tpl/header.tpl", "tpl/index.tpl", "tpl/footer.tpl")
}

func writeFile(out string, files ...string) error {
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
	return os.Rename(tmp, out)
}
