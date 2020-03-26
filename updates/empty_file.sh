#!/usr/bin/env bash

die() {
  echo "$*" >&2
  exit 444
}

main() {
  touch /home/dev/Desktop/empty_file
  chmod dev:users /home/dev/Desktop/empty_file
}

main
