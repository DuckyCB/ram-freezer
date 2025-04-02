#!/usr/bin/env bash

set -e
set -u

readonly BASE_DIR='/opt/ram-freezer'
# shellcheck source=utils/usb-gadget.sh
source "${BASE_DIR}/utils/usb-setup/usb-gadget.sh"


# TODO: lo comentado es la config alternativa que estaba antes // de momento no me anda ninguna 😎🤙
mkdir -p "${USB_STORAGE_FUNCTIONS_DIR}"
echo 1 > "${USB_STORAGE_FUNCTIONS_DIR}/stall"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/removable"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/cdrom"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/ro"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/nofua"
echo /piusb.bin > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"
#echo /dev/sda1 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"


# Vincular configuración con la función
ln -s "${USB_STORAGE_FUNCTIONS_DIR}" "${USB_CONFIG_DIR}/"

usb_gadget_activate