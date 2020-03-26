#!/usr/bin/env bash

die() {
  echo "$*" >&2
  exit 444
}

main() {
  # opens firefox.
  firefox
}

main
