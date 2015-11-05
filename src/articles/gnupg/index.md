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

## Generate a new key

	gpg --gen-key

## List recipients of a encrypted message

	gpg --list-packets < file.gpg

## Display Fingerprint

	gpg --fingerprint <key_name> | grep "Key fingerprint" | tr -s " " | cut -d ' ' -f 5-

## Encrypt to self by default

Add `default-recipient-self` to `$HOME/.gnupg/gpg.conf`
