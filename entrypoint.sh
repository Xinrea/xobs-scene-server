#!/bin/bash
if [ -n "$1" ]; then
  echo $1
  sed -i "s/ENDPOINT = \".*\"/ENDPOINT = \"$1\"/" static/main/index.html
fi
./xobs-scene-server
