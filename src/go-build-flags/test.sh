#!/bin/sh
set -e

f=$(mktemp /tmp/go-build-flags-XXXX)
trap "rm -f $f" EXIT

if go version | grep "go version go1.4" > /dev/null; then
  go build -ldflags "-X main.BUILD_INFO $(HISTORY_LIMIT=10 bash ./build_info.sh)" -o $f
else
	go build -ldflags "-X main.BUILD_INFO=$(HISTORY_LIMIT=10 bash ./build_info.sh)" -o $f
fi
$f
