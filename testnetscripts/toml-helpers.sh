#!/bin/bash

function toml_set(){
  echo not working yet
  exit 2
  FILE=${1:?"TOML file required"}
  SECTION=${2:?"SECTION required"}
  KEY=${3:?"KEY required"}
  VALUE=${4:?"VALUE required"}
  sed -i "s/(\[$SECTION\][^\[]\{0,\}\s\{0,\})\#\{0,\}\s\{0,\}$KEY\s\{0,\}=.\{0,\}/\1\n$KEY = $VALUE/"
}

function toml_get(){
  echo not working yet
  exit 2
  FILE=${1:?"TOML file required"}
  SECTION=${2:?"SECTION required"}
  KEY=${3:?"KEY required"}
  VALUE=${4:?"VALUE required"}
  cat $FILE | sed -n "s/\[$SECTION\][^\[]\{0,\}\s\{0,\}\#\{0,\}\s\{0,\}$KEY\s\{0,\}=\s\{0,\}(.\{0,\})\n/\1/p"
  grep -e"\[$SECTION\][^\[]\{0,\}\s\{0,\}\#\{0,\}\s\{0,\}$KEY\s\{0,\}=\s\{0,\}(.\{0,\})\n"
}