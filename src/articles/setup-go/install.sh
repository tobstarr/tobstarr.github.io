#!/bin/sh
set -e

VERSION=${VERSION-:1.5.1}
ARCH=$(uname | awk '{print tolower($0)}')-amd64
DST=/usr/local/go${VERSION}

sudo_prefix=""

if [[ $(id -u) != "0" ]]; then
	sudo_prefix="sudo"
fi

if [[ ! -e $DST/bin/go ]]; then
	if ! which curl; then
		apt-get install -y curl
	fi

	TMP=$(mktemp -d /tmp/goXXXXXX)
	cd ${TMP} && curl -SfL https://storage.googleapis.com/golang/go${VERSION}.${ARCH}.tar.gz | tar xfz - --strip=1

	$sudo_prefix mv ${TMP} ${DST}
	$sudo_prefix chmod 0755 ${DST}
fi

rm -Rf $HOME/pkg

$sudo_prefix ln -nfs ${DST} /usr/local/go

# it is also helpful to delete $GOPATH/pkg (at least when upgrading go < 1.5)
# also make sure to add /usr/local/go/bin to your PATH
