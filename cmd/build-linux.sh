#!/bin/bash

export GOOS=linux

cmds=$(ls -l ./ | awk '/^d/ {print $NF}')

for d in $cmds; do
  cd $d || exit
  echo "building ${d} ..."
  go build -tags=jsoniter .
  cd ../ || exit
done

[ -f "cmd.zip" ] && mv cmd.zip cmd.zip.bak

for d in $cmds; do
  zip cmd.zip -jg "./${d}/${d}"
done

[ -f cmd.zip.bak ] && rm -rf cmd.zip.bak
