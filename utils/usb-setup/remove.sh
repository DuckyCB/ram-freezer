#!/usr/bin/env bash

set -e
set -u


printf "Eliminando usb-gadget...\n"

source "/opt/ram-freezer/utils/usb-setup/usb-gadget.sh"


# Service
systemctl stop usb-gadget.service
systemctl disable usb-gadget.service
rm /etc/systemd/system/usb-gadget.service
rm /usr/lib/systemd/system/usb-gadget.service
systemctl daemon-reload

# Gadget
bash /opt/ram-freezer/utils/usb-setup/remove-usb-gadget.sh


printf "usb-gadget eliminado\n"
