package gotime

import (
	"testing"
	"time"
)

func TestFromUnit(t *testing.T) {
	tests := []struct {
		Input  int64
		Unit   time.Duration
		Output string
	}{
		{1450345359, time.Second, "2015-12-17T10:42:39"},
		{1450345359123, time.Millisecond, "2015-12-17T10:42:39.123"},
		{1450345359123456, time.Microsecond, "2015-12-17T10:42:39.123456"},
		{1450345359123456789, time.Nanosecond, "2015-12-17T10:42:39.123456789"},
	}

	const format = "2006-01-02T15:04:05.999999999"

	for _, tst := range tests {
		v := TimeFromUnit(tst.Input, tst.Unit).Format(format)
		if v != tst.Output {
			t.Errorf("expected to get %q, was %q", tst.Output, v)
		}
	}

}

func TestFromMillis(t *testing.T) {
	const millis = 1450345359123
	tm := timeFromMillis(millis)
	if v, ex := tm.Format("2006-01-02T15:04:05.000"), "2015-12-17T10:42:39.123"; ex != v {
		t.Errorf(`expected tm.Format("2006-01-02T15:04:05.000") to be %q, was %q`, ex, v)
	}
}
