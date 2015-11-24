#!/bin/sh
set -e

REPO_PATH=$HOME/src/github.com/tobstarr/tobstarr.github.io/.git
GIT_COMMIT=$(git --git-dir $REPO_PATH rev-parse HEAD)

HISTORY_LIMIT=${HISTORY_LIMIT:-50}
GIT_HISTORY=""
for name in $(git --git-dir $REPO_PATH rev-list HEAD -n $HISTORY_LIMIT); do
	n=$(echo $name | cut -b 1-12)
	if [[ -n $GIT_HISTORY ]]; then
		GIT_HISTORY=${GIT_HISTORY},
	fi
	GIT_HISTORY=${GIT_HISTORY}'"'$n'"'
done

if [[ -n $GIT_HISTORY ]]; then
	GIT_HISTORY='['$GIT_HISTORY']'
fi

BUILT_AT=$(TZ=utc date +"%Y-%m-%dT%H:%M:%S")

build=$(cat <<EOF
{
  "build_at":    "$BUILT_AT",
  "git_history": $GIT_HISTORY,
  "git_sha":     "$(echo $GIT_COMMIT | cut -b 1-12)",
  "changes": $(if git status --porcelain > /dev/null 2>&1; then echo true; else false; fi),
  "hostname": "$(hostname)",
  "user": "$(whoami)"
}
EOF
)

echo $build | jq -c -r '.'
