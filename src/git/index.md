# Git

## Rebase first commit

	git rebase --root -i


## list branchs of a commit

	git --git-dir=$GIT_DIR branch --contains $REF -q -r | awk '{ print $1 }' | xargs

## find out if a commit "contains" another commit

	git rev-list f5ee3a8 | grep $(git rev-parse b9281ca)


## empty pushes

	git commit --allow-empty -m "update" && git push aws

## update hook

	$1: old commit
	$2: new commit
	$3: symbolic reference
