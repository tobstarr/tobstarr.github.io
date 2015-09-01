# Setup OSX

## Homebrew

### defaults

* elasticsearch
* gnu-sed
* gnu-tar
* gnu-time
* gnupg
* imagemagick
* jq
* keychain
* mysql
* pv
* redis
* the_silver_searcher
* tmux
* vim
* zsh

### phrase

* icu4c
* qt5


## Symlinks

	export DOTFILES=$HOME/src/github.com/tobstarr/tobstarr.github.io/dotfiles

	ln -nfs ${DOTFILES}/vimrc ~/.vimrc
	ln -nfs ${DOTFILES}/tmux.conf ~/.tmux.conf
	ln -nfs ${DOTFILES}/gitconfig ~/.gitconfig
