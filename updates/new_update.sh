#!/usr/bin/env bash

die() {
  echo "$*" >&2
  exit 444
}

main() {
  # opens firefox.
  sudo -u dev DISPLAY=:0 DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus firefox
}

main
