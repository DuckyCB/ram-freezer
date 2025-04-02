#!/usr/bin/env bash

set -e
set -u

source "/opt/ram-freezer/utils/usb-setup/usb-gadget.sh"

# TODO: se podria usar un almacenamiento emulado si /dev/sda1 no existe
#dd if=/dev/zero of=/piusb.bin bs=1M count=1024
#mkdosfs /piusb.bin -F 32 -I

mkdir -p "${USB_STORAGE_FUNCTIONS_DIR}"
echo 1 > "${USB_STORAGE_FUNCTIONS_DIR}/stall"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/removable"
echo /dev/sda1 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/cdrom"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/ro"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/nofua"
#echo /piusb.bin > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"


ln -s "${USB_STORAGE_FUNCTIONS_DIR}" "${USB_CONFIG_DIR}/"

printf "Vault configurado\n"
