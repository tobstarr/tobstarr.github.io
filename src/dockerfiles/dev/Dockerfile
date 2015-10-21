FROM ubuntu:15.04

RUN apt-get update && apt-get upgrade -y && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install --no-install-recommends -y tmux zsh jq ca-certificates bzr mercurial python python-dev libncurses5-dev curl bzip2 build-essential git && rm -rf /var/lib/apt/lists/*

RUN DIR=$(mktemp -d /tmp/vimXXXXXX) && \
	cd $DIR && \
	curl -SsfL ftp://ftp.vim.org/pub/vim/unix/vim-7.4.tar.bz2 | tar xfj - --strip=1 && \
	./configure --enable-pythoninterp --with-features=huge --with-python-config-dir=$(/usr/bin/python-config --configdir) && \
	make && \
	make install

RUN mkdir -p $HOME/.vim.swp

RUN mkdir -p $HOME/.vim && git clone https://github.com/gmarik/Vundle.vim.git ${HOME}/.vim/bundle/Vundle.vim

RUN mkdir -p $HOME/src/github.com/tobstarr && \
	git clone https://github.com/tobstarr/tobstarr.github.io.git $HOME/src/github.com/tobstarr/tobstarr.github.io

RUN ln -nfs $HOME/src/github.com/tobstarr/tobstarr.github.io/dotfiles/tmux.conf $HOME/.tmux.conf && \
	ln -nfs $HOME/src/github.com/tobstarr/tobstarr.github.io/dotfiles/vimrc.conf $HOME/.vimrc && \
	ln -nfs $HOME/src/github.com/tobstarr/tobstarr.github.io/dotfiles/gitconfig $HOME/.gitconfig && \
	ln -nfs $HOME/src/github.com/tobstarr/tobstarr.github.io/dotfiles/vim-snippets UltiSnips

ADD zshrc /root/.zshrc 

RUN usermod -s /usr/bin/zsh root

ENTRYPOINT ["/usr/bin/tmux"]
