#!/bin/bash
set -e # abort in error

GEMINABOX_VERSION=0.12.4
RUBY_VERSION=2.2.3
TAG=geminabox:${GEMINABOX_VERSION}
NAME=geminabox
BUILD_DIR=$(mktemp -d /tmp/docker-build-XXXXXX)
PORT=8888
cd $BUILD_DIR

# delete build dir when done
trap "rm -Rf $BUILD_DIR" EXIT

cat > Dockerfile <<EOF
FROM ruby:${RUBY_VERSION}

RUN gem install puma --no-ri --no-rdoc
RUN gem install geminabox -v ${GEMINABOX_VERSION} --no-ri --no-rdoc

RUN mkdir -p /app

COPY config.ru /app/config.ru
COPY entrypoint.sh /app/entrypoint.sh

RUN chmod a+x /app/entrypoint.sh

WORKDIR /app

ENV RUBYGEMS_PROXY true

ENTRYPOINT ["/bin/bash", "/app/entrypoint.sh"]
EOF

cat > config.ru <<EOF
require "rubygems"
require "geminabox"

Geminabox.data = "/data"
run Geminabox::Server
EOF

cat > entrypoint.sh <<"EOF"
puma -p $PORT -b tcp://0.0.0.0
EOF

docker build -t ${TAG} .

# delete old geminabox container if exists
if docker inspect --type container ${NAME} > /dev/null 2>&1; then
	docker rm -f ${NAME} > /dev/null
fi

# start new geminabox container
# store data in mounted host /data/docker/geminabox to "survive" restarts
container_id=$(docker run -d -v /data/docker/geminabox:/data -e PORT=${PORT} -p ${PORT}:${PORT} --name ${NAME} ${TAG})

cat <<EOF
Now please add this to your local bundler config at $HOME/.bundle/config:
---
BUNDLE_MIRROR__HTTPS://RUBYGEMS__ORG/: http://127.0.0.1:${PORT}
BUNDLE_MIRROR__HTTP://RUBYGEMS__ORG/: http://127.0.0.1:${PORT}
EOF
