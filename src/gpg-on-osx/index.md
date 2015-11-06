# GnuPG on OSX

## Setup

	`brew install gpg-agent keychain pinentry-mac`

	# $HOME/.gnupg/gpg-agent.conf
	# add this line, create file if it does not exist
	pinentry-program /usr/local/bin/pinentry-mac

	# $HOME/.gnupg/gpg.conf
	# make sure this line exists in that file, uncomment or add
	use-agent

	# $HOME/.profile, $HOME/.zshrc
	# add this line to your profile file
	eval $(keychain --agents gpg --eval)
