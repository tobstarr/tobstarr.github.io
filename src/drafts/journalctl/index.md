# journalctl

## Examples for --since and --until

* -10s
* -1d
* yesterday

## Flags

flag |explanation
---|---
`-u systemd-logind` | filter all messages by a given service |
`-x` | log lines with explanation |
`-e` | jump to the end of the pager
`-f` | follow journal
`-c` | start logs from a given cursor position
`--show-cursor` | display cursor at end of the logs
`--utc` | show timestamps in utc

## Write to yournal

`systemd-cat -t uptime uptime`

## Parse Journal Timestamp in Go

	func parseTimestamp(in string) (time.Time, error) {
		// raw is the unix nanon
		unixNano, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return timeFromUnixNano(unixNano), nil
	}

	func timeFromUnixNano(in int64) time.Time {
		f := float64(in) / 1000000
		sec := int64(f)
		rem := f - float64(sec)
		return time.Unix(int64(sec), int64(1000000000*rem))
	}
