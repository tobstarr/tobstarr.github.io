#!/bin/bash
set -e # abort on error

TAG=geminabox:${GEMINABOX_VERSION}
NAME=geminabox
PORT=8888

# Setup docker build directory
BUILD_DIR=$(mktemp -d /tmp/docker-build-XXXXXX)
cd $BUILD_DIR

# delete build dir when done
trap "rm -Rf $BUILD_DIR" EXIT

# build docker images
docker build -t geminabox:current .

# delete old geminabox container if exists
if docker inspect --type container ${NAME} > /dev/null 2>&1; then
	docker rm -f ${NAME} > /dev/null
fi

# start new geminabox container
# store data in volume /data/docker/geminabox to "survive" restarts
container_id=$(docker run -d -v /data/docker/geminabox:/data -e PORT=${PORT} -p ${PORT}:${PORT} --name ${NAME} geminabox:current)

cat <<EOF
Now please add this to your local bundler config at $HOME/.bundle/config:
---
BUNDLE_MIRROR__HTTPS://RUBYGEMS__ORG/: http://127.0.0.1:${PORT}
BUNDLE_MIRROR__HTTP://RUBYGEMS__ORG/: http://127.0.0.1:${PORT}
EOF
