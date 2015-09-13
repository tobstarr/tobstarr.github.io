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
