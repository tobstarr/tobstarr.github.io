package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	b, err := readMessagePayload()
	if err != nil {
		return err
	} else if len(b) == 0 {
		return fmt.Errorf("unable to find message")
	}
	dbg.Printf("logging message text: %q", b)
	m, err := NewMessage()
	if err != nil {
		return err
	}
	m.Message = string(b)
	s := &Scratcher{}
	return s.Write(m)
}

func readMessagePayload() ([]byte, error) {
	var reader io.Reader
	if len(os.Args) < 2 {
		reader = os.Stdin
	} else {
		reader = strings.NewReader(strings.Join(os.Args[1:], " "))
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return bytes.TrimSpace(b), nil
}

type Message struct {
	Time      time.Time `json:"time,omitempty"`
	MachineID string    `json:"machine_id,omitempty"`
	Hostname  string    `json:"hostname,omitempty"`
	Message   string    `json:"message,omitempty"`
}

func NewMessage() (h *Message, err error) {
	h = &Message{Time: time.Now()}
	h.MachineID, err = readMachineID()
	if err != nil {
		return nil, err
	}
	h.Hostname, err = os.Hostname()
	if err != nil {
		return nil, err
	}
	return h, nil
}

type Scratcher struct {
	Root string
}

func (s *Scratcher) CurrentPath(m *Message) string {
	root := s.Root
	if root == "" {
		root = os.ExpandEnv("$HOME/.scratch/logs")
	}
	suffix := []string{}
	if m.Hostname != "" {
		suffix = append(suffix, m.Hostname)
	}
	if m.MachineID != "" {
		suffix = append(suffix, m.MachineID)
	}
	suffix = append(suffix, ".log")
	return filepath.Join(root, m.Time.Format("2006/01/02/2006-01-02T15")+strings.Join(suffix, "."))
}

func (s *Scratcher) Write(m *Message) error {
	pc := s.CurrentPath(m)
	if err := os.MkdirAll(filepath.Dir(pc), 0755); err != nil {
		return err
	}
	dbg.Printf("writing to %s", pc)
	// open the file for the current hour, create if necessary, append if exists, exclusively (so we do not mess things up)

	var f *os.File

	var err error
	for _, flag := range []int{os.O_CREATE | os.O_RDWR | os.O_EXCL, os.O_APPEND | os.O_RDWR | os.O_EXCL} {
		f, err = os.OpenFile(pc, flag, 0644)
		if err != nil {
			continue
		}
		defer f.Close()
		break
	}
	if err != nil {
		return err
	}
	return json.NewEncoder(f).Encode(m)
}

func readMachineID() (string, error) {
	// ubuntu sets a unique machine id for each installation, not sure about other linuxes
	b, err := ioutil.ReadFile("/etc/machine-id")
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(b)), nil
}

func debugStream() io.Writer {
	if os.Getenv("DEBUG") == "true" {
		return os.Stderr
	}
	return ioutil.Discard
}

var dbg = log.New(debugStream(), "[DEBUG] ", log.Lshortfile)
