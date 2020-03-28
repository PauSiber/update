#!/usr/bin/env bash

GO_VERSION="1.14"
WORKING_DIRECTORY="/home/dev/.update/"

die() {
  echo "$*" >&2
  exit 444
}

clear_screen() {
  clear
  printf "\e[3J"
}

do_you_confirm() {
  while :
  do
    echo ""
    echo ""
    echo "Installing Update System..."
    sleep 1
    echo ""
    echo -e "Update system needs go environment."
    echo ""
    echo ""
    echo -en "$1 Do you confirm? [y/n] "
    read answer
    echo ""
    case "$answer" in
      [yY] ) echo -e "Installation was started.\nPlease do not cancel the operation.\nOtherwise your system can be broke." && break ;;
      [nN] ) die "Okay. Update service is cancelled." && break ;;
      *    ) clear_screen && echo "(!) Please use only Y or N" ;;
    esac
  done
}

push_enter() {
  echo ""
  echo -e "Installation was completed."
  echo -e $1
  echo ""
  echo -e "[ Push enter to continue ] "
  read answer
}

install_go() {
  # Installs GVM if it is not.
  if [[ ! -s "/home/dev/.gvm/" ]]; then
    do_you_confirm "Gvm will be installed."
    zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer) >/dev/null 2>&1
    source /home/dev/.gvm/scripts/gvm
    push_enter "`gvm version`"
    clear_screen
  fi

  source /home/dev/.gvm/scripts/gvm
  # Installs go if it is not.
  if [[ ! -s "/home/dev/.gvm/gos/go${GO_VERSION}/bin/go" ]]; then
    do_you_confirm "Go ${GO_VERSION} will be installed."
    gvm install go${GO_VERSION} -B >/dev/null 2>&1
    gvm use go${GO_VERSION} --default >/dev/null 2>&1
    push_enter "`go version`"
    clear_screen
  fi
}

run_go() {
  go run main.go
}

main() {
  cd ${WORKING_DIRECTORY}
  install_go
  run_go
}

main
