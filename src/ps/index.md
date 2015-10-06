# PS

## Check for systemd

	if [[ $(ps -p 1 -o comm -h) = "systemd" ]]; then echo "it is systemd"; fi

## Check for upstart
	
	if [[ $(ps -p 1 -o comm -h) = "init" ]]; then echo "it is init"; fi
