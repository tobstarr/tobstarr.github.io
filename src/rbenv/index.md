# rbenv

    git clone https://github.com/sstephenson/rbenv.git ~/.rbenv
    git clone https://github.com/sstephenson/ruby-build.git ~/.rbenv/plugins/ruby-build

    if [ -e $HOME/.rbenv ]; then
      export PATH=$HOME/.rbenv/shims:${PATH}:${HOME}/.rbenv/bin
      eval "$(rbenv init -)"
    fi
