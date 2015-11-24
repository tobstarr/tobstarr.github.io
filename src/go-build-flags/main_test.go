package main

import (
	"encoding/json"
	"os/exec"
	"testing"
	"time"
)

func TestLoadBuildInfo(t *testing.T) {
	b, err := exec.Command("bash", "./test.sh").CombinedOutput()
	if err != nil {
		t.Fatalf("%s: %s", err, b)
	}

	var info *BuildInfo
	if err := json.Unmarshal(b, &info); err != nil {
		t.Fatal(err)
	}

	if v, ex := time.Since(info.BuiltAt), 10*time.Second; v > ex {
		t.Errorf("expected time since %s to be < %s, was %s", info.BuiltAt.String(), ex, v)
	}
}
