#!/bin/bash
set -xe

VERSION=0.12.4
RUBY_VERSION=2.2.3
TAG=geminabox:${VERSION}
NAME=geminabox
BUILD_DIR=$(mktemp -d /tmp/docker-build-XXXXXX)
cd $BUILD_DIR

# delete build dir when done
trap "rm -Rf $BUILD_DIR" EXIT

cat > Dockerfile <<EOF
FROM ruby:${RUBY_VERSION}

RUN gem install puma --no-ri --no-rdoc
RUN gem install geminabox -v ${VERSION} --no-ri --no-rdoc

RUN mkdir -p /app

COPY config.ru /app/config.ru
COPY entrypoint.sh /app/entrypoint.sh

RUN chmod a+x /app/entrypoint.sh

WORKDIR /app

ENV RUBYGEMS_PROXY true
ENV PORT 8888
EXPOSE 8888

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
	docker rm -f ${NAME}
fi

# start new geminabox container
# store data in mounted host /data/docker/geminabox to "survive" restarts
docker run -d -v /data/docker/geminabox:/data -e PORT=8888 -p 8888:8888 --name ${NAME} ${TAG}
