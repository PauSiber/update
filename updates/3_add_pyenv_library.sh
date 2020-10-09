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
  git clone https://github.com/pyenv/pyenv.git /home/dev/.pyenv
  echo 'export PYENV_ROOT="$HOME/.pyenv"' >> /home/dev/.zshrc
  echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> /home/dev/.zshrc
  echo "Added pyenv library to system."
  success "(✔) You can use pyenv command to switch between python versions."
}

cd ${WORKING_DIRECTORY}
main
