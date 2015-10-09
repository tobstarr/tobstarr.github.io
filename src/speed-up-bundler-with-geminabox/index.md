# Speed up bundler with Geminabox

Waiting for `bundle install` can be looooooooong. A lot of time seems to be network latency.

[Geminabox](https://rubygems.org/gems/geminabox) can help you save quite a bit of time with not much setup overhead.

## Setup

You could just install geminabox via `gem install geminabox` but I prefer to run it inside a docker container. This is how:

Download: [setup_geminabox.sh](setup_geminabox.sh)

{{ require "src/speed-up-bundler-with-geminabox/geminabox.sh" | code}}

This script

1. builds a new docker container
2. removes an old geminabox container if exists
3. starts a new geminabox container in the background

See to find out how you can [make docker run on your machine](https://docs.docker.com/).

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
