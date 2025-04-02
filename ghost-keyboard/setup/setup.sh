#!/usr/bin/env bash

set -e
set -u


# Kernel modules
readonly MODULES_PATH='/etc/modules'
# TODO: chequear si es este path que rompe el teclado
CONFIG_FILE=/boot/firmware/config.txt

if [[ ! -e $CONFIG_FILE ]]; then
    CONFIG_FILE=/boot/config.txt
fi

if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi

if ! grep --quiet '^dtoverlay=dwc2$' "${CONFIG_FILE}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${CONFIG_FILE}"
fi

# Setup ghost keyboard
cd /opt/ram-freezer/ghost-keyboard/setup

chmod +x init-usb-gadget.sh
./init-usb-gadget.sh

# USB gadget service
cp usb-gadget.service /lib/systemd/system/usb-gadget.service

systemctl daemon-reload
systemctl enable usb-gadget.service

# Exec permissions
chmod +x remove.sh
chmod +x remove-usb-gadget.sh
