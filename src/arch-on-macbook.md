# Arch on Macbook

https://wiki.archlinux.org/index.php/MacBook#OS_X_with_Arch_Linux

mkdir -p /boot/efi

mount /dev/sda1 /boot/efi

grub-install

grub-mkconfig -o /boot/grub/grub.cfg}

pacman -S grub efibootmgr dosfstools


# mount -t efivarfs efivarfs /sys/firmware/efi/efivars

https://wiki.archlinux.de/title/Systemverschl%C3%BCsselung_mit_dm-crypt

https://wiki.archlinux.de/title/UEFI_Installation

https://wiki.archlinux.org/index.php/dm-crypt/Encrypting_an_entire_system#Encrypted_boot_partition_.28GRUB.29

GRUB_CMDLINE_LINUX="... cryptdevice=UUID=<device-UUID>:lvm root=/dev/mapper/MyStorage-rootvol ..."
GRUB_ENABLE_CRYPTODISK=y

grub-mkconfig -o /boot/grub/grub.cfg
grub-install --target=x86_64-efi --efi-directory=/boot/efi --bootloader-id=grub --recheck
