#!/bin/bash
set -e

IMAGES_ROOT=/opt/rkt/images
RUBY_VERSION=2.2.3
GEMINABOX_VERSION=0.12.4
PUMA_VERSION=2.15.3
DEFAULT_PORT=8888

function image_path() {
  name=$1
  version=$2
  if [[ -z $name || -z $version ]]; then
    cat "both name and version must be present" > /dev/stderr
    exit 1
  fi
  image_name=library-${name}-${version}.aci
  image_path=$IMAGES_ROOT/$image_name
  if [[ ! -f $image_path ]]; then
    dir=$(basename $image_path)
    mkdir -p $dir
    pushd $dir > /dev/null
    cat "downloading ${name}:${version}" > /dev/stderr
    /usr/local/bin/docker2aci docker://${name}:${version}
    popd > /dev/null
  fi
  echo $image_path
}

function write_files() {
cat > entrypoint.sh <<"EOF"
#!/bin/bash

if [[ -z $BIND ]]; then
  echo "BIND (e.g. tcp://127.0.0.1:8888) must be set"
  exit 1
fi

exec puma -b $BIND
EOF

cat > config.ru <<"EOF"
require "rubygems"
require "geminabox"

Geminabox.data = "/data"
run Geminabox::Server
EOF
}

d=$(mktemp -d /tmp/build-XXXX)
pushd $d

# write files
write_files

# exec commands
set -xe
acbuild begin $(image_path ruby $RUBY_VERSION)
acbuild copy entrypoint.sh /app/entrypoint.sh
acbuild copy config.ru /app/config.ru
acbuild run -- apt-get update
acbuild run -- apt-get upgrade -y
acbuild run -- rm -rf var/lib/apt/lists
acbuild run -- gem install puma -v $PUMA_VERSION
acbuild run -- gem install geminabox -v $GEMINABOX_VERSION
acbuild run -- mkdir -p /app
acbuild run -- chmod a+x /app/entrypoint.sh
acbuild environment add RUBYGEMS_PROXY true
acbuild environment add BIND tcp://127.0.0.1:$DEFAULT_PORT
acbuild port add http tcp $DEFAULT_PORT
acbuild mount add data /data
acbuild set-exec /app/entrypoint.sh
acbuild set-working-directory /app
acbuild set-name geminabox
acbuild write $IMAGES_ROOT/geminabox-0.12.4.aci
acbuild end
rm -Rf $d

# run via
# mkdir /opt/rkt/geminabox
# rkt run /opt/rkt/images/geminabox-0.12.4.aci --net=host --insecure-options=all --volume data,kind=host,source=/opt/rkt/geminabox
