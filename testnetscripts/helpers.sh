# checking for required commands.
ensure_command_exists () {
  if ! command -v $1 &> /dev/null
  then
    echo "$1 could not be found"
    exit
  fi
} 

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

