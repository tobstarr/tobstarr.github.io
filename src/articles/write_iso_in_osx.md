# Write ISO in OSX

Source: http://www.ubuntu.com/download/desktop/create-a-usb-stick-on-mac-osx

	hdiutil convert -format UDRW -o ~/path/to/target ~/path/to/ubuntu.iso # no extension required
	diskutil list
	diskutil unmountDisk /dev/diskN
	sudo dd if=/path/to/downloaded.img of=/dev/rdiskN bs=1m
