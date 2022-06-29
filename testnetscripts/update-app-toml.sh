#!/bin/bash

CONFIG_FILE_NAME="config/app.toml"

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

function enable_rest() {
  TRG=${1:?"Target folder required"}

  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "api" "enable" "true"
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "api" "swagger" "true"
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "api" "enabled-unsafe-cors" "true"
}

function update_app_toml(){
  SRC=${1:?"Source folder required"}
  TRG=${2:?"Target folder required"}

  enable_rest "$TRG"
  toml_set "${TRG%/}/$CONFIG_FILE_NAME" "grpc-web" "enabled-unsafe-cors" "true"
}