#!/bin/sh
set -xe

VERSION=1.5
ARCH="linux-amd64" # use darwin-amd64 for OSX
DST=/usr/local/go${VERSION}

TMP=$(mktemp -d /tmp/goXXXXXX)
cd ${TMP} && curl -SfL https://storage.googleapis.com/golang/go${VERSION}.${ARCH}.tar.gz | tar xfz - --strip=1

sudo mv ${TMP} ${DST}
sudo ln -nfs ${DST} /usr/local/go

# it is also helpful to delete $GOPATH/pkg (at least when upgrading go < 1.5)
# also make sure to add /usr/local/go/bin to your PATH
