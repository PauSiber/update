#!/usr/bin/env bash

WORKING_DIRECTORY="/home/dev/.update/"

die() {
  echo "$*" >&2
  exit 444
}

working_directory() {
  if [[ ! -s ${WORKING_DIRECTORY} ]]; then
    mkdir -p ${WORKING_DIRECTORY}
  fi
  cd ${WORKING_DIRECTORY}
}

main() {
  if [[ ! -s .git ]]; then
    git clone https://github.com/boratanrikulu/update.git .
  else
    git pull
  fi
  clear

  if [[ ! -s "/home/dev/.local/share/systemd/user/update_dev.service" ]]; then
    mkdir -p /home/dev/.local/share/systemd/user
    cp update_dev.service /home/dev/.local/share/systemd/user/update_dev.service
    # If update_dev.service is not created yet,
    # Then this script must run by root user.
    # Change permission for system files.
    chown dev:users -R /home/dev/.local/share/systemd/
    chown dev:users -R ${WORKING_DIRECTORY}
    sudo -u dev XDG_RUNTIME_DIR=/run/user/1000 systemctl --user enable update_dev.service || die "Error eccur while enabling service."
    systemctl disable update_dev.service
    sudo -u dev XDG_RUNTIME_DIR=/run/user/1000 systemctl --user start update_dev.service
    die "Service was enabled. Reboot is needed."
  fi

  if [[ -s "go.sh" ]]; then
    chmod u+x go.sh
  else
    die "There is no go installation script."
  fi

  sudo -u dev DISPLAY=:0 DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus xfce4-terminal \
       --title "Update - PauSiber Dev" \
       --hide-menubar \
       --hide-toolbar \
       -e "./go.sh" \
       --hold
}

working_directory
main
