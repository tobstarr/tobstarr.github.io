# Git

## Rebase first commit

	git rebase --root -i


## list branchs of a commit

	git --git-dir=$GIT_DIR branch --contains $REF -q -r | awk '{ print $1 }' | xargs

## find out if a commit "contains" another commit

	git rev-list f5ee3a8 | grep $(git rev-parse b9281ca)
