#!/usr/bin/env bash

GO_VERSION="1.14"
WORKING_DIRECTORY="/home/dev/Desktop/update"

die() {
  echo "$*" >&2
  exit 444
}

install_go() {
  if [[ ! -s "/home/dev/.gvm/" ]]; then
    while :
    do
      echo ""
      echo -en "Gvm will be installed. Do you confirm? [y/n] "
      read answer
      echo ""
      case "$answer" in
        [yY] ) echo -e "Installation was started.\nPlease do not cancel the operation.\nOtherwise you system can be broke." && break ;;
        [nN] ) die "Okay. Update service is cancelled." && break ;;
        *    ) clear && echo "(!) Please use only Y or N" ;;
      esac
    done
    zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer) >/dev/null 2>&1
    source /home/dev/.gvm/scripts/gvm
    echo -e "Installation was completed."
    echo -e `gvm version`
    echo -e "[enter]"
    read answer
    clear
  fi

  if [[ ! -e "/home/dev/.gvm/gos/go${GO_VERSION}/bin/go" ]]; then
    while :
    do
      echo ""
      echo -en "Go ${GO_VERSION} will be installed. Do you confirm? [y/n] "
      read answer
      echo ""
      case "$answer" in
        [yY] ) echo -e "Installation was started.\nPlease do not cancel the operation.\nOtherwise you system can be broke." && break ;;
        [nN] ) die "Okay. Update service is cancelled." && break ;;
        *    ) clear && echo "(!) Please use only Y or N" ;;
      esac
    done
    gvm install go${GO_VERSION} -B >/dev/null 2>&1
    gvm use go${GO_VERSION} --default >/dev/null 2>&1
    echo -e "Installation was completed."
    echo -e `go version`
    echo -e "[enter]"
    read answer
    clear
  fi

  source /home/dev/.gvm/scripts/gvm
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
