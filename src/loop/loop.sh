#!/bin/zsh

SLEEP=${SLEEP:-5}

while true; do
  res=$($@)
  clear
  date
  echo $res
  sleep $SLEEP
done

