#!/bin/bash


if [ -z "$1" ]; then
  echo "provide the path to the private key"
  exit 1
fi

if [ ! -f "$1" ]; then
  echo "the given file $1 does not exist."
  exit 1
fi

SCRIPT_LOCATION=$( dirname -- "${BASH_SOURCE[0]}")
# CAUTION do not change this, without extending the .gitignore file to avoid publishing your private key
KEY_NAME="ssh_key"
if [ -f "$SCRIPT_LOCATION/$KEY_NAME" ]; then
  echo "Apparently a private key is already linked."
  exit 1
fi

ln -s $1 $SCRIPT_LOCATION/$KEY_NAME