#!/usr/bin/env bash

set -e
set -u

source "/opt/ram-freezer/utils/usb-setup/usb-gadget.sh"

# TODO: se podria usar un almacenamiento emulado si /dev/sda1 no existe
#dd if=/dev/zero of=/piusb.bin bs=1M count=1024
#mkdosfs /piusb.bin -F 32 -I
USB_PATH="/dev/sda1"

mkdir -p "${USB_STORAGE_FUNCTIONS_DIR}"
echo 1 > "${USB_STORAGE_FUNCTIONS_DIR}/stall"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/removable"
echo "${USB_PATH}" > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/cdrom"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/ro"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/nofua"
#echo /piusb.bin > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"


ln -s "${USB_STORAGE_FUNCTIONS_DIR}" "${USB_CONFIG_DIR}/"

# Storage device label
# Para FAT32 o exFATg
dosfslabel "${USB_PATH}" "${USB_STORAGE_DEVICE_NAME}"

printf "Vault configurado\n"
