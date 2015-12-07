#!/bin/sh
set -xe

# fixed in https://github.com/appc/acbuild/issues/132

export d=$(mktemp -d /tmp/acbuild-XXXXX)
trap "rm -Rf $d" EXIT
export GOPATH=$d
export dir=$GOPATH/src/github.com/appc/acbuild
mkdir -p $dir

pushd $dir

curl -sL https://github.com/appc/acbuild/archive/0c8fb80a67955ff51ee151b46e10818a759ee9e7.tar.gz | tar xz --strip=1

go get -d github.com/appc/acbuild/acbuild
go build -o $GOPATH/bin/acbuild github.com/appc/acbuild/acbuild
sudo mv $GOPATH/bin/acbuild /usr/local/bin/acbuild
