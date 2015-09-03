# GnuPG

## Export Public Key

	gpg --output tobstarr.gpg --armor --export tobstarr@gmail.com

## Import Public Key (e.g. from keybase)

	echo "some key" | gpg --import // e.g. from keybase

## Public Key from keybase

	curl -Ls https://keybase.io/tobstarr/key.asc | gpg --import

## List Keys

	gpg -k

## Encrypt a stream without confirmation

	cat some.file | gpg -e -a --trust-model always --recipient tobstarr@gmail.com
