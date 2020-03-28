## PauSiber Dev - Update Service

[![Go Report Card](https://goreportcard.com/badge/github.com/pausiber/update)](https://goreportcard.com/report/github.com/pausiber/update)

> An update service that brings updates from PauSiber team to users.  


You will be up to date.  
You will be getting updates from PauSiber team..

## Features

- Included gvm and go installation script.
- Checking updates on every boots.
- A tiny interface to use.

## Installation

**Hey!**  
If you have already [**PauSiber Dev**](https://dev.pausiber.xyz/).  
Do not worry.

You already have the update system.
Run this command to get magic.
```shell
systemctl start update_dev.service
```

## Usage

[![usage-video](usage.jpg)](https://www.youtube.com/watch?v=VgaM_Ejru6o)

## How to send updates?

It is simple.  
Add update information to [**updates.json**](./updates/updates.json).  
```json
{
  "authority": "PauSiber Community",
  "name": "PauSiber Dev",
  "version": "2.0",
  "updates": [
    {},
    {},
    {
      "id": LAST ID + 1,
      "name": "YOUR UPDATE NAME",
      "description": "YOUR UPDATE DESCRIPTION",
      "fileName": "YOUR UPDATE FILE NAME",
      "publishTime": "YOUR UPDATE PUBLISH TIME"
    }
  ]
}
```
Add your update script to [**updates/**](./updates) folder.
Like this one:
```shell
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
  echo "Added conio library to system."
  success "You can use getch and getche methods now."
}

cd ${WORKING_DIRECTORY}
main
```

## To-Do

- [ ] Add "do all updates" (--no-confirm) to cli.
- [ ] Add notifications for Pacman and Aur updates.
