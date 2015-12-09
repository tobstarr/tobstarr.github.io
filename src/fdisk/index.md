# fdisk

	fdisk -l /dev/xvdm 2>&1 | grep "doesn.t contain"; } && echo ';' | sfdisk /dev/xvdm && mkfs.ext4 /dev/xvdm1; } || /bin/true
