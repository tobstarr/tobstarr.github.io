#!/bin/bash
set -e

NAME=kubectl
VERSION=${VERSION:-1.3.6}
DL_PATH=/tmp/${NAME}.${VERSION}
URL=https://storage.googleapis.com/kubernetes-release/release/v${VERSION}/bin/linux/amd64/${NAME}

echo "downloading $URL"

curl -SfL $URL > $DL_PATH
trap "rm -f $DL_PATH" EXIT
install $DL_PATH /usr/local/bin/${NAME}
