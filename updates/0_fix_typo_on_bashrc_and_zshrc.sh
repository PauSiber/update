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
  sed -i 's+opn+open+g' .zshrc
  sed -i 's+opn+open+g' .bashrc
  sed -i 's+boratanrikulu.me+boratanrikulu.dev+g' .zshrc
  sed -i 's+boratanrikulu.me+boratanrikulu.dev+g' .bashrc
  echo "alias open='xdg-open' "  >> /home/dev/.zshrc
  echo "alias open='xdg-open' "  >> /home/dev/.bashrc
  chown dev:users .zshrc
  chown dev:users .bashrc
  success "(✔) Fixed typos on zshrc and bashrc."
}

cd ${WORKING_DIRECTORY}
main
