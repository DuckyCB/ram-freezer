#!/usr/bin/env bash

# Echo commands to stdout.
set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

# Enable the dwc2 kernel driver, which we need to emulate USB devices with USB OTG.
readonly MODULES_PATH='/etc/modules'
if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi
readonly BOOT_CONFIG_PATH='/boot/config.txt'
if ! grep --quiet '^dtoverlay=dwc2$' "${BOOT_CONFIG_PATH}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${BOOT_CONFIG_PATH}"
fi

cd /opt/ram-freezer/ghost-keyboard/setup

chmod +x init-usb-gadget
./init-usb-gadget

# Start keyboard at boot
cp usb-gadget.service /lib/systemd/system/usb-gadget.service

systemctl daemon-reload
systemctl enable usb-gadget.service
