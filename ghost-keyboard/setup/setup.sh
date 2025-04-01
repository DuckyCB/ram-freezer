#!/usr/bin/env bash

# Echo commands to stdout.
set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

readonly MODULES_PATH='/etc/modules'
# TODO: chequear si es este path que rompe el teclado
CONFIG_FILE=/boot/firmware/config.txt

# check if $CONFIG_FILE exists or go back to old path
if [[ ! -e $CONFIG_FILE ]]; then
    CONFIG_FILE=/boot/config.txt
fi

# Enable the dwc2 kernel driver, which we need to emulate USB devices with USB OTG.
if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi

if ! grep --quiet '^dtoverlay=dwc2$' "${CONFIG_FILE}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${CONFIG_FILE}"
fi

cd /opt/ram-freezer/ghost-keyboard/setup

chmod +x init-usb-gadget
./init-usb-gadget

# Start USB Gadget at boot
cp usb-gadget.service /lib/systemd/system/usb-gadget.service

systemctl daemon-reload
systemctl enable usb-gadget.service
