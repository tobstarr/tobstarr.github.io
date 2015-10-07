# Docker

## Install debian packages

	RUN apt-get update && apt-get install --no-install-recommends -y curl && rm -rf /var/lib/apt/lists/*

## Download URL with curl

	RUN mkdir -p /dl && cd /dl && curl -SsfLO http://nginx.org/download/nginx-1.8.0.tar.gz

## Install go

	RUN mkdir -p /go && curl -SsfL https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar xfz - --strip=1
	ENV GOROOT /go
	ENV GOPATH /

## Validate Docker Running

	for i in $(seq 1 10); do
		if docker ps > /dev/null 2>&1; then
			exit 0
		else
			echo "not running => waiting"
			sleep 1
		fi
	done
	exit 1

## Touch Gemfile Gemfile.lock

	touch -m -t 200601021504.05 Gemfile Gemfile.lock bundle_config

## Run Registry

	docker run -p 5000:5000 --rm --name registry -v $HOME/docker-distribution.yml:/etc/docker/registry/config.yml registry:2.1.1

with config

	version: 0.1
	log:
	  fields:
		service: registry
	http:
	  addr: :5000
	  headers:
		X-Content-Type-Options: [nosniff]
	storage:
	  cache:
		blobdescriptor: inmemory
	  filesystem:
		rootdirectory: /var/lib/registry
	health:
	  storagedriver:
		enabled: true
		interval: 10s
		threshold: 3
