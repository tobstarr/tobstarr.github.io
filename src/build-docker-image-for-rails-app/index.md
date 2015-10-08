# Build Docker image for your Rails application

## Build

### Base Image

	FROM ruby:2.2.3

### Update packages

    RUN apt-get update && \
        DEBIAN_FRONTEND=noninteractive apt-get upgrade -y \
        rm -rf /var/lib/apt/lists/*

We use `rm -rf /var/lib/apt/lists` to make the image footprint smaller

### Install required packages (for your rails app)

    RUN apt-get update && \
        DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y libxml2-dev imagemagick unzip libmysqlclient-dev && \
        rm -rf /var/lib/apt/lists/*

By default we update all sources before we install packages.
