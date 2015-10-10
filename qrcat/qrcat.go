package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"

	"code.google.com/p/rsc/qr"
)

func main() {
	c := make(chan struct{})
	s := httptest.NewServer(qrHandler(c, os.Stdin))
	defer s.Close()
	if err := open(s.URL); err != nil {
		log.Fatal(err)
	}
	<-c
}

func qrHandler(c chan struct{}, in io.Reader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer close(c)
		err := func() error {
			b, err := ioutil.ReadAll(in)
			if err != nil {
				return err
			} else if len(b) == 0 {
				return fmt.Errorf("no text provided")
			}
			code, err := qr.Encode(string(b), qr.L)
			if err != nil {
				return err
			} else {
				w.Header().Set("Content-Type", "image/png")
				w.Write(code.PNG())
			}
			return nil
		}()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

func open(res string) error {
	for _, cmd := range []string{"xdg-open", "open"} {
		if _, err := exec.LookPath(cmd); err == nil {
			return exec.Command(cmd, res).Run()
		}
	}
	return fmt.Errorf("no command found to open URLs")
}
