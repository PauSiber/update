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
  cp ./.update/updates/_files/1_conio.h /usr/include/conio.h
  success "Added conio library to system.\nYou can use getch and getche methods now."
}

cd ${WORKING_DIRECTORY}
main
