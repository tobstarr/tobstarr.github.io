package main

import (
	"testing"
	"time"
)

func TestTimeToMillis(t *testing.T) {
	tim := time.Date(1970, 1, 1, 0, 0, 0, int(10*time.Millisecond), time.UTC)
	if v, ex := timeToMillis(tim), int64(10); ex != v {
		t.Errorf("expected time to millis to be %d, was %d", ex, v)
	}
}
