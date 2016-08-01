#!/bin/bash
set -xe

CHANNEL=${CHANNEL:-beta}
DIR=$HOME/coreos/$CHANNEL

sudo dnf install qemu-system-x86 qemu-img

mkdir -p $DIR
pushd $DIR > /dev/null

if [[ ! -f coreos_production_qemu.sh || ! -f coreos_production_qemu.img ]]; then
	for name in coreos_production_qemu.sh coreos_production_qemu.sh.sig coreos_production_qemu_image.img.bz2 coreos_production_qemu_image.img.bz2.sig; do
		curl -SfLOz ${name} https://${CHANNEL}.release.core-os.net/amd64-usr/current/${name}
	done
	curl https://coreos.com/security/image-signing-key/CoreOS_Image_Signing_Key.asc | gpg --import
	gpg --verify coreos_production_qemu.sh.sig
	gpg --verify coreos_production_qemu_image.img.bz2.sig
	bzip2 -d coreos_production_qemu_image.img.bz2
	chmod +x coreos_production_qemu.sh
fi

./coreos_production_qemu.sh -a ~/.ssh/id_rsa.pub -- -nographic
