#!/usr/bin/env bash

# Project
export USB_SETUP_PATH="/opt/ram-freezer/utils/usb-setup"
readonly USB_SETUP_PATH
export KEYBOARD_SETUP_PATH="/opt/ram-freezer/ghost-keyboard/setup"
readonly KEYBOARD_SETUP_PATH
export STORAGE_SETUP_PATH="/opt/ram-freezer/vault/setup"
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
export USB_STORAGE_DEVICE_NAME="USB_Vault"
readonly USB_STORAGE_DEVICE_NAME
export USB_STORAGE_FUNCTIONS_DIR="functions/${USB_STORAGE_NAME}"
readonly USB_STORAGE_FUNCTIONS_DIR

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
	# Check if /sys/class/udc is empty
	if [ -z "$(ls /sys/class/udc)" ]; then
		echo "No UDC found. Exiting."
		echo "Please check if the kernel module is loaded and the device is connected."
		exit 1
	fi

	ls /sys/class/udc >"${USB_DEVICE_PATH}/UDC"
	chmod 777 /dev/hidg0
}
export -f usb_gadget_activate

function usb_gadget_deactivate {
	echo '' >"${USB_DEVICE_PATH}/UDC"
}
export -f usb_gadget_deactivate
