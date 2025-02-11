#!/usr/bin/env bash

# Echo commands to stdout.
set -x

# Exit on first error.
set -e

# Treat undefined environment variables as errors.
set -u

mkdir /opt/umkey
cd /opt/umkey

# Enable the dwc2 kernel driver, which we need to emulate USB devices with USB OTG.
readonly MODULES_PATH='/etc/modules'
if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi
readonly BOOT_CONFIG_PATH='/boot/config.txt'
if ! grep --quiet '^dtoverlay=dwc2$' "${BOOT_CONFIG_PATH}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${BOOT_CONFIG_PATH}"
fi

ENABLE_RPI_HID_PATH=/opt/enable-rpi-hid

#cp /home/user/enable-rpi-hid "$ENABLE_RPI_HID_PATH"
#chmod +x "$ENABLE_RPI_HID_PATH"

# Start keyboard at boot
cp usb-gadget.service /lib/systemd/system/usb-gadget.service

systemctl daemon-reload
systemctl enable usb-gadget.service
