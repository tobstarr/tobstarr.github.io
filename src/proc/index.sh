# Index

## meminfo

  cat /proc/meminfo | grep "MemTotal\|MemFree\|Buffers\|Cache\|Active\|SwapTotal\|SwapFree\|Committed_AS" | sed 's/ kB//'g | tr -s "\n" ";"  | sed 's/ //g'

## procs

  for dir in $(find /proc -maxdepth 1 -mindepth 1 -type d); do
    if [[ -f $dir/stat ]]; then
      echo "cmdline=\"$(strings $dir/cmdline | xargs)\" stat=\"$(cat $dir/stat)\" statm=\"$(cat $dir/statm)\"" > /dev/null
    fi
  done
