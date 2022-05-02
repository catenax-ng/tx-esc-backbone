#!/bin/bash

./build.sh

if [ ! -d $HOME/.esc-backbone/ ]; then
  starport chain init
fi

docker-compose up