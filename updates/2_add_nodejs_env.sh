#!/usr/bin/env bash

WORKING_DIRECTORY="/home/dev"

die() {
  echo "$*" >&2
  exit 444
}

success() {
  echo "$*" >&2
  exit 0
}

main() {
  pacman -S --noconfirm nodejs npm
  echo "NodeJS and npm installed to system."
  success "You can use NodeJS on your terminal."
}

cd ${WORKING_DIRECTORY}
main
