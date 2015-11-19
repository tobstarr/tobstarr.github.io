FROM ubuntu:15.10

RUN apt-get update && apt-get upgrade -y && rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install --no-install-recommends -y git jq ca-certificates silversearcher-ag && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /root/src/github.com/tobstarr && git clone https://github.com/tobstarr/tobstarr.github.io /root/src/github.com/tobstarr/tobstarr.github.io

RUN bash /root/src/github.com/tobstarr/tobstarr.github.io/setup_vim.sh
RUN VERSION=1.4.3 bash /root/src/github.com/tobstarr/tobstarr.github.io/setup_go.sh
RUN bash /root/src/github.com/tobstarr/tobstarr.github.io/setup_vim.sh
RUN ln -nfs /root/src/github.com/tobstarr/tobstarr.github.io/dotfiles/vimrc /root/.vimrc

ENV GOPATH=$HOME \
    GOROOT=/usr/local/go

RUN /usr/local/bin/vim +BundleInstall +qall > /dev/null 2>&1
RUN /usr/local/bin/vim +GoInstallBinaries +qall > /dev/null 2>&1

ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:$HOME/bin:$GOROOT/bin

ENTRYPOINT ["/usr/local/bin/vim"]
