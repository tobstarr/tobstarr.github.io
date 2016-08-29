#!/bin/bash
set -e

VERSION=${VERSION:-1.3.6}
DL_PATH=/tmp/kubelet.${VERSION}
URL=https://storage.googleapis.com/kubernetes-release/release/v${VERSION}/bin/linux/amd64/kubelet

echo "downloading $URL"

curl -SfL $URL > $DL_PATH
install $DL_PATH /usr/local/bin/
