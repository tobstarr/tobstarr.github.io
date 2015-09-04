# Docker

## Start registry

	docker run -p 5000:5000 --rm --name registry -v $HOME/secrets/docker-distribution.yml:/etc/docker/registry/config.yml registry:2.1.1
