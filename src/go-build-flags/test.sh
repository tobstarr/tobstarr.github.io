#!/bin/sh
set -e

f=$(mktemp /tmp/go-build-flags-XXXX)
trap "rm -f $f" EXIT

go build -ldflags "-X main.REVISION $(HISTORY_LIMIT=10 bash ./build_info.sh)" -o $f
$f
