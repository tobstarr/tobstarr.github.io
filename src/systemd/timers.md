# Systemd Timers

## Service
	#/etc/systemd/system/journal-backup.service
	[Unit]
	Description=Backup journalctl

	[Service]
	ExecStart=/home/tobias/bin/ts-journal

## Timer
	#/etc/systemd/system/journal-backup.timer
	[Unit]
	Description=Journal Backup Service

	[Timer]
	OnCalendar=*:0/5:0 # every 5 minutes

	[Install]
	WantedBy=timers.target

## Display Timer Status

	systemctl list-timers journal-backup.timer

## List Logs
	
	journalctl -u journal-backup

## Reenable timer

	systemctl reenable --now journal-backup.timer


## Enable

	systemctl enable journal-backup.timer
	systemctl start journal-backup.timer

## OnCalendar Patterns

name | pattern
---|---
every 15 seconds | OnCalendar=*:*:0/15 
every 10 minutes | OnCalendar=*:0/10:0
