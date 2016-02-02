# Bash

## Set variables with default

	FOO=${VARIABLE:-default}

## pipe and print program output at the same time
	
	echo 1 | tee >(logger -i -t hello)

## shell dimensions

	tput cols
	tput lines
