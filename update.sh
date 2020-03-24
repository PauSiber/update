#!/usr/bin/env bash

WORKING_DIRECTORY="/home/dev/Desktop/update"

die() {
  echo "$*" >&2
  exit 444
}

main() {
  if [[ ! -s "/home/dev/.local/share/systemd/user/update_dev.service" ]]; then
    mkdir -p /home/dev/.local/share/systemd/user
    cp update_dev.service /home/dev/.local/share/systemd/user/update_dev.service
    systemctl --user enable update_dev.service
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

cd ${WORKING_DIRECTORY}
main
