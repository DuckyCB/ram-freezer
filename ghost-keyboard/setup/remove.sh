#!/usr/bin/env bash

set -e
set -x
set -u


printf "Eliminando usb-gadget..."

# Service
systemctl stop usb-gadget.service
systemctl disable usb-gadget.service
rm /etc/systemd/system/usb-gadget.service
rm /usr/lib/systemd/system/usb-gadget.service
systemctl daemon-reload

/opt/ram-freezer/ghost-keyboard/setup/remove-usb-gadget.sh

printf "usb-gadget eliminado"