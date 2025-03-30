#!/usr/bin/env bash

# Echo commands to stdout.
set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

MODULES_PATH='/etc/modules'
CONFIG_FILE=/boot/firmware/config.txt
CMDLINE_FILE=/boot/firmware/cmdline.txt

# check if $CONFIG_FILE exists or go back to old path
if [[ ! -e $CONFIG_FILE ]]; then
    CONFIG_FILE=/boot/config.txt
    CMDLINE_FILE=/boot/cmdline.txt
fi

# Enable the dwc2 kernel driver, which we need to emulate USB devices with USB OTG.
if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi

if ! grep --quiet '^dtoverlay=dwc2$' "${BOOT_CONFIG_PATH}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${BOOT_CONFIG_PATH}"
fi

cd /opt/ram-freezer/ghost-keyboard/setup

chmod +x init-usb-gadget
./init-usb-gadget

# Start USB Gadget at boot
cp usb-gadget.service /lib/systemd/system/usb-gadget.service

systemctl daemon-reload
systemctl enable usb-gadget.service
