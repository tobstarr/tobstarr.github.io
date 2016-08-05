# Fedora

## Fonts
https://www.reddit.com/r/Fedora/comments/3o6ijr/trying_fedora_23_for_the_first_time_why_do_fonts/:

## Volume
https://faq.i3wm.org/question/125/how-to-change-the-systems-volume.1.html

## Wifi

	wicd-curses
		or
	nmtui
		or
	nm-applet


## Mouse and Touchpad Values

	xset q

## Installed packages

	dnf list installed

## Wifi Driver

	https://gist.githubusercontent.com/onpubcom/7f41dc9cbe90556b2113/raw/a69939c941319741744bea28dadf273f118d67a2/fedora23_broadcom_wl_install.sh

	https://gist.github.com/jamespamplin/7a803fd5be61d4f93e0c5dcdea3f99ee

## dist-upgrade

	dnf upgrade
	dnf install dnf-plugin-system-upgrade
	dnf system-upgrade download --releasever=24
	dnf system-upgrade reboot

## rpmfusion
	dnf install --nogpgcheck \
  http://download1.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm  \
  http://download1.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
