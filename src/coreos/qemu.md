# qemu

	dnf install qemu-system-x86 qemu-img
	mkdir coreos; cd coreos
	wget https://stable.release.core-os.net/amd64-usr/current/coreos_production_qemu.sh
	wget https://stable.release.core-os.net/amd64-usr/current/coreos_production_qemu.sh.sig
	wget https://stable.release.core-os.net/amd64-usr/current/coreos_production_qemu_image.img.bz2
	wget https://stable.release.core-os.net/amd64-usr/current/coreos_production_qemu_image.img.bz2.sig
	gpg --verify coreos_production_qemu.sh.sig
	gpg --verify coreos_production_qemu_image.img.bz2.sig
	bzip2 -d coreos_production_qemu_image.img.bz2
	chmod +x coreos_production_qemu.sh

## Start

	./coreos_production_qemu.sh -a ~/.ssh/id_rsa.pub -- -nographic

## Teminate

	ctrl+alt+q
