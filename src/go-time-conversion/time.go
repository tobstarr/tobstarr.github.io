package gotime

import "time"

func timeFromMillis(in int64) time.Time {
	return TimeFromUnit(in, time.Millisecond)
}

func TimeFromUnit(in int64, unit time.Duration) time.Time {
	sec := in / int64(time.Second/unit)
	reminder := in % int64(time.Second/unit)
	nano := reminder * int64(unit/time.Nanosecond)
	return time.Unix(sec, nano)
}
