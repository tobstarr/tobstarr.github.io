package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var BUILD_INFO string

func main() {
	i, err := loadBuildInfo()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", b)
}

type BuildInfo struct {
	BuiltAt    time.Time `json:"built_at,omitempty"`
	Changes    bool      `json:"changes,omitempty"`
	GOARCH     string    `json:"go_arch,omitempty"`
	GOOS       string    `json:"go_os,omitempty"`
	GOVERSION  string    `json:"go_version,omitempty"`
	GitHistory []string  `json:"git_history,omitempty"`
	GitSHA     string    `json:"git_sha,omitempty"`
	Hostname   string    `json:"hostname,omitempty"`
	User       string    `json:"user,omitempty"`
}

func loadBuildInfo() (i *BuildInfo, err error) {
	if err := json.Unmarshal([]byte(BUILD_INFO), &i); err != nil {
		return nil, err
	}
	i.GOOS = runtime.GOOS
	i.GOARCH = runtime.GOARCH
	i.GOVERSION = runtime.Version()
	return i, nil
}
