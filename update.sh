#!/usr/bin/env bash

WORKING_DIRECTORY="/home/dev/.update/"
GIT_REPO_ADDRESS="https://github.com/pausiber/update.git"

die() {
  echo "$*" >&2
  exit 444
}

working_directory() {
  # Creates working directory if it do not exist.
  if [[ ! -s ${WORKING_DIRECTORY} ]]; then
    mkdir -p ${WORKING_DIRECTORY}
  fi
  cd ${WORKING_DIRECTORY}
}

clear_screen() {
  clear
  printf "\e[3J"
}

main() {
  if [[ ! -s .git ]]; then
    # Clones the repo.
    git clone ${GIT_REPO_ADDRESS} .
  else
    # Pulls it if the repo is already cloned.
    git pull
  fi
  clear_screen

  # If update_dev.service is not created yet then this script must run by root user.
  # Disables root's service. Creates a new service for the user (--user).
  if [[ ! -s "/home/dev/.local/share/systemd/user/update_dev.service" ]]; then
    mkdir -p /home/dev/.local/share/systemd/user
    cp update_dev.service /home/dev/.local/share/systemd/user/update_dev.service
    chown dev:users -R /home/dev/.local/share/systemd/
    # Changes permission for system files.
    chown dev:users -R ${WORKING_DIRECTORY}
    sudo -u dev XDG_RUNTIME_DIR=/run/user/1000 systemctl --user enable update_dev.service \
                                                        || die "Error eccur while enabling service."
    systemctl disable update_dev.service
    sudo -u dev XDG_RUNTIME_DIR=/run/user/1000 systemctl --user start update_dev.service
    echo "Service was enabled. Reboot is needed."
    exit 0
  fi

  # Gives executable permission to user to run "go.sh".
  if [[ -s "go.sh" ]]; then
    chmod u+x go.sh
  else
    die "There is no go installation script."
  fi

  # Opens the terminal to the user's screen (display 0).
  # Executes "go.sh".
  sudo -u dev DISPLAY=:0 DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus xfce4-terminal \
       --title "Update - PauSiber Dev" \
       --hide-menubar \
       --hide-toolbar \
       --geometry 100x25+40+100 \
       -e "./go.sh" \
       --hold
}

working_directory
main
