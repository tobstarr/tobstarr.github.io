# Systemd

## Reload Services

	sudo systemctl daemon-reload


# Env

	Environment="ONE=one" 'TWO=two two'
	EnvironmentFile=/etc/default/service

# User

	[Service]
	User=tobias # run as specific user
