# Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
#
# See the NOTICE file(s) distributed with this work for additional
# information regarding copyright ownership.
#
# This program and the accompanying materials are made available under the
# terms of the Apache License, Version 2.0 which is available at
# https://www.apache.org/licenses/LICENSE-2.0.
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.
#
# SPDX-License-Identifier: Apache-2.0

# checking for required commands.
ensure_command_exists () {
  if ! command -v $1 &> /dev/null
  then
    echo "$1 could not be found"
    exit
  fi
} 

ensure_command_exists git
ensure_command_exists basename

is_unsigned_int () {
  case $1 in 
    ''|*[!0-9]*) echo "$1 not an unsigned integer" >&2; return 1;;
    *) return 0;;
  esac
}


apply_on_each (){
  array=($@)
  func=$1
  params=("${array[@]:1}")
  for entry in "${params[@]}"
  do
    $func "$entry"
  done
}


home_name () {
   echo "$(basename $1)"
}

create_a_local_empty_repo () {
  GIT_DISCOVERY_ACROSS_FILESYSTEM=1
  local _GIT_REPO=${1:?"Provide the folder for the local repo"}
  local _REPO_BRANCH=${2:-"main"}
  git init --bare --shared $_GIT_REPO
  cd $_GIT_REPO
  git symbolic-ref HEAD refs/heads/"$_REPO_BRANCH"
  cd -

}

