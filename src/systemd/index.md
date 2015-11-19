# Systemd

## Reload Services

	sudo systemctl daemon-reload


# ENV

	Environment="ONE=one" 'TWO=two two'
	EnvironmentFile=/etc/default/service

# User

	[Service]
	User=tobias # run as specific user

## Run commands in systemd context

	sudo systemd-run bash -c "while true; do date; sleep 10; done"
