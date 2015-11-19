#!/bin/bash
set -xe

sudo_prefix=""

if [[ $USER != "" && $USER != "root" ]]; then
	sudo_prefix="sudo"
fi

updated=false
for p in python python-dev libncurses5-dev curl bzip2 build-essential git; do
  if ! dpkg -s $p >/dev/null 2>&1; then
    if [[ $updated == false ]]; then
      updated=true
      $sudo_prefix apt-get update
    fi
    $sudo_prefix apt-get install -y $p
  else
    echo "pkg $p already installed"
  fi
done


if [[ ! -e /usr/local/bin/vim ]]; then
  echo "downloading and extracting go"
  tmp=$(mktemp -d /tmp/vimXXXXXX)
  cd $tmp

  curl -SsfL ftp://ftp.vim.org/pub/vim/unix/vim-7.4.tar.bz2 | tar xfj - --strip=1
  ./configure --enable-pythoninterp --with-features=huge --with-python-config-dir=$(/usr/bin/python-config --configdir)
  make
  $sudo_prefix make install
  cd $HOME
  rm -Rf $tmp
else
  echo "go already installed"
fi

if [[ ! -e ${HOME}/.vim/bundle/Vundle.vim ]]; then
	git clone https://github.com/gmarik/Vundle.vim.git ${HOME}/.vim/bundle/Vundle.vim
fi

# /usr/local/bin/vim +BundleInstall +qall > /dev/null 2>&1
# /usr/local/bin/vim +GoInstallBinaries +qall 2>&1
