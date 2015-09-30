# Docker and iptables

Exposing container ports on a host where the firewall is managed with iptables can lead to unwanted issues if you do not plan to expose those ports to the world also.

If the firewall is setup before docker, running a container like this

	`docker run -p 5432:5432 postgres`

opens up port `5432` to the world, even though your iptables configuration should forbid this.

## Solution

The solution I currently use is to expose those containers to a specific network interface only. 

If I only need to __access this container from the host__ running docker itself I use

	`docker run -p 127.0.0.1:5432:5432 postgres`

If I also want to __access this container from other containers__ on the same host I bind it to the docker network interface (defaults to 172.17.42.1)

	`docker run -p 172.17.42.1:5432:5432 postgres`

## Conclusion

There are probably much better and more robust solutions to this very problem but I am not an iptables expert so for the time being binding to dedicated interfaces is good enough for me. Feel free to ping me via [@tobstarr](https://twitter.com/tobstarr) if you know a better solution.
