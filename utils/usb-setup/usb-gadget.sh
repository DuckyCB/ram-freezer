#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh


# Project
export BASE_PATH="/opt/ram-freezer"
readonly BASE_PATH
export USB_SETUP_PATH="${BASE_PATH}/utils/usb-setup"
readonly USB_SETUP_PATH
export KEYBOARD_SETUP_PATH="${BASE_PATH}/ghost-keyboard/setup"
readonly KEYBOARD_SETUP_PATH
export STORAGE_SETUP_PATH="${BASE_PATH}/vault/setup"
readonly STORAGE_SETUP_PATH

# Kernel
export MODULES_PATH="/etc/modules"
readonly MODULES_PATH
export CONFIG_FILE=/boot/firmware/config.txt
if [[ ! -e $CONFIG_FILE ]]; then
    CONFIG_FILE=/boot/config.txt
fi
## Systemd
export SYSTEM_SERVICES="/lib/systemd/system"
readonly SYSTEM_SERVICES

# Gadget
export USB_DEVICE_DIR="ram-freezer"
readonly USB_DEVICE_DIR
export USB_GADGET_PATH="/sys/kernel/config/usb_gadget"
readonly USB_GADGET_PATH
export USB_DEVICE_PATH="${USB_GADGET_PATH}/${USB_DEVICE_DIR}"
readonly USB_DEVICE_PATH

export USB_STRINGS_DIR="strings/0x409"
readonly USB_STRINGS_DIR

## Keyboard
export USB_KEYBOARD_FUNCTIONS_DIR="functions/hid.keyboard"
readonly USB_KEYBOARD_FUNCTIONS_DIR
## Storage
export USB_STORAGE_NAME="mass_storage.usb0"
readonly USB_STORAGE_NAME
export USB_STORAGE_DEVICE_NAME="USB_VAULT"
readonly USB_STORAGE_DEVICE_NAME
export USB_STORAGE_FUNCTIONS_DIR="functions/${USB_STORAGE_NAME}"
readonly USB_STORAGE_FUNCTIONS_DIR
export LOCAL_STORAGE_FILE="${BASE_PATH}/bin/piusb.bin"
readonly LOCAL_STORAGE_FILE
export LOCAL_STORAGE_SIZE="18432"
readonly LOCAL_STORAGE_SIZE
export USB_DRIVE_PATH="/dev/sda1"
readonly USB_DRIVE_PATH

## USB config
export USB_CONFIG_INDEX=1
readonly USB_CONFIG_INDEX
export USB_CONFIG_DIR="configs/c.${USB_CONFIG_INDEX}"
readonly USB_CONFIG_DIR
export USB_ALL_CONFIGS_DIR="configs/*"
readonly USB_ALL_CONFIGS_DIR
export USB_ALL_FUNCTIONS_DIR="functions/*"
readonly USB_ALL_FUNCTIONS_DIR

# Functions
function usb_gadget_activate {
	if [ -z "$(ls /sys/class/udc)" ]; then
		log_error "No UDC found. Exiting. Please check if the kernel module is loaded and the device is connected."
	fi

	ls /sys/class/udc >"${USB_DEVICE_PATH}/UDC"
	chmod 777 /dev/hidg0
}
export -f usb_gadget_activate

function usb_gadget_deactivate {
	echo '' >"${USB_DEVICE_PATH}/UDC"
}
export -f usb_gadget_deactivate
