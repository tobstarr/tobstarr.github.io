# Speed up bundler with Geminabox

Waiting for `bundle install` to finish can be pretty annoying, especially if you work on e.g. a large rails project that uses a few gems.

__TODO__: add proof and write about the reasons

[Geminabox](https://rubygems.org/gems/geminabox) can help you save quite a bit of time with not much setup overhead.

## Setup
You could just install geminabox via `gem install geminabox` but I prefer to run it inside a docker container. This is how:

### Build image

To build the geminabox container you need three files, the `Dockerfile`, a `config.ru` file and `entrypoint.sh`. It is best to place those files in a new and empty directory. You can use `dir=$(mktemp -d /tmp/build-XXX) && cd $dir` to create and cd into a new temporary build directory.

__# Dockerfile__

{{ require "src/speed-up-bundler-with-geminabox/Dockerfile" | code}}

__# config.ru__

{{ require "src/speed-up-bundler-with-geminabox/config.ru" | code}}

__# entrypoint.sh__

{{ require "src/speed-up-bundler-with-geminabox/entrypoint.sh" | code}}

Download: __TODO__ add download link to tar archive

Now you can create a new geminabox image by running `docker build -t geminabox:current .`

See to find out how you can [make docker run on your machine](https://docs.docker.com/).

### Run

	# delete old geminabox container if exists
	docker rm -f geminabox > /dev/null

	# start new geminabox container and daemonize (-d) on
	# port 8888 (-e PORT=8888) and expose that port to the same port on the current host (-p 8888:8888)
	# store data in volume /data/docker/geminabox (-v) to "survive" restarts
	docker run -d -v /data/docker/geminabox:/data -e PORT=8888 -p 8888:8888 --name geminabox geminabox:current

## Test

Geminabox should now be running on port 8888 on your docker host. You can test it by opening http://DOCKER_HOST_IP:8888.

## Bundler setup

The only thing left is to the geminabox mirror to your bundler configuration at `$HOME/.bundle/config`:

	---
	BUNDLE_MIRROR__HTTPS://RUBYGEMS__ORG/: http://DOCKER_HOST_IP:${PORT}
	BUNDLE_MIRROR__HTTP://RUBYGEMS__ORG/: http://DOCKER_HOST_IP:${PORT}


## Docker

Caching gems is especially helpful when building docker images which use bundler. 

Here is an example for how you can use geminabox to also proxy bundler when building docker images.


### Dockerfile

	FROM ruby:2.2.3
	RUN mkdir -p /app
	ADD Gemfile.lock /app
	ADD Gemfile /app
	ADD bundle_config /root/.bundle/config
	RUN cd /app && bundle


### bundle_config
	---
	BUNDLE_MIRROR__HTTPS://RUBYGEMS__ORG/: http://172.17.42.1:${PORT}
	BUNDLE_MIRROR__HTTP://RUBYGEMS__ORG/: http://172.17.42.1:${PORT}

This relies on docker being bound to the default interface `172.17.42.1`, adapt if necessary.


### Gemfile

	source "https://rubygems.org"
	gem "rails"


## Debug

You can always debug what is going on with geminabox by running `docker logs -f geminabox`.
