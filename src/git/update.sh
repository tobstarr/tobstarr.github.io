#!/bin/bash
set -e

refname="$1"
oldrev="$2"
newrev="$3"

wd=$(pwd)
d=$(mktemp -d /tmp/docker-test-XXXX)
trap "rm -Rf $d" EXIT

pushd $d

git --git-dir=$wd archive $newrev | tar x

GIT_COMMIT=$newref REFNAME=$refname bash ./git_hook.sh
