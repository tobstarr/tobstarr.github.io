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
	OnBootSec=1m
	OnUnitActiveSec=5m

	[Install]
	WantedBy=timers.target

## Display Timer Status

	systemctl list-timers journal-backup.timer

## List Logs
	
	journalctl -u journal-backup

