#!/bin/bash
set -xe

which git > /dev/null 2>&1 || sudo apt-get install -y git

CONF_DIR=$HOME/src/github.com/tobstarr/tobstarr.github.io

if [[ ! -d $CONF_DIR/.git ]]; then
        rm -Rf $CONF_DIR
        mkdir -p $(dirname $CONF_DIR)
        git clone https://github.com/tobstarr/tobstarr.github.io $CONF_DIR
fi

rm -Rf $HOME/.vim $HOME/.vim.swp $HOME/.vimrc $HOME/.gitconfig $HOME/.tmux.conf
mkdir -p $HOME/.vim $HOME/.vim.swp

if [[ ! -f /usr/local/bin/vim ]]; then
        bash $CONF_DIR/src/articles/setup-vim/setup_vim.sh
fi

if [[ ! -f /usr/local/go/bin/go ]]; then
        bash $CONF_DIR/src/articles/setup-go/install.sh
fi

ln -nfs $CONF_DIR/dotfiles/vimrc $HOME/.vimrc
ln -nfs $CONF_DIR/dotfiles/vim-snippets $HOME/.vim/UltiSnips
ln -nfs $CONF_DIR/dotfiles/tmux.conf $HOME/.tmux.conf
ln -nfs $CONF_DIR/dotfiles/gitconfig $HOME/.gitconfig

cat <<EOF > $HOME/.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME
export PATH=$GOROOT/bin:$GOPATH/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games
alias vi=$(which vim)
EOF
