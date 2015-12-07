test -f $HOME/.profile && source $HOME/.profile

ZSH=$HOME/.oh-my-zsh

if [[ ! -e $ZSH ]]; then
	git clone git://github.com/robbyrussell/oh-my-zsh.git ${HOME}/.oh-my-zsh
fi

CASE_SENSITIVE="true"
COMPLETION_WAITING_DOTS="true"
DISABLE_AUTO_UPDATE="true"

cdpath=$(. ~)
source $ZSH/oh-my-zsh.sh

export PS1="%n@%m:%4c $ "

if [ -n $TMUX ]; then
  export TERM="screen-256color"
else
  export TERM="xterm-256color"
fi
