#!/bin/sh
set -xe

VERSION=${VERSION-:1.5.1}
ARCH=$(uname | awk '{print tolower($0)}')-amd64
DST=/usr/local/go${VERSION}

echo "using VERSION=${VERSION}"

if [[ ! -e $DST/bin/go ]]; then
	if ! which curl; then
		apt-get install -y curl
	fi

	TMP=$(mktemp -d /tmp/goXXXXXX)
	cd ${TMP} && curl -SfL https://storage.googleapis.com/golang/go${VERSION}.${ARCH}.tar.gz | tar xfz - --strip=1

	sudo mv ${TMP} ${DST}
	sudo chmod 0755 ${DST}
fi

rm -Rf $HOME/pkg

sudo ln -nfs ${DST} /usr/local/go

# it is also helpful to delete $GOPATH/pkg (at least when upgrading go < 1.5)
# also make sure to add /usr/local/go/bin to your PATH
