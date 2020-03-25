#!/usr/bin/env bash

die() {
  echo "$*" >&2
  exit 444
}

main() {
  echo "This operation is run by" `whoami`
}

main
