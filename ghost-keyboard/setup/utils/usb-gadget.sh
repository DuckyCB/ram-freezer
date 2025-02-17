#!/bin/bash

export USB_DEVICE_DIR="ghost-keyboard"
readonly USB_DEVICE_DIR
export USB_GADGET_PATH="/sys/kernel/config/usb_gadget"
readonly USB_GADGET_PATH
export USB_DEVICE_PATH="${USB_GADGET_PATH}/${USB_DEVICE_DIR}"
readonly USB_DEVICE_PATH

export USB_STRINGS_DIR="strings/0x409"
readonly USB_STRINGS_DIR
export USB_KEYBOARD_FUNCTIONS_DIR="functions/hid.keyboard"
readonly USB_KEYBOARD_FUNCTIONS_DIR

export USB_CONFIG_INDEX=1
readonly USB_CONFIG_INDEX
export USB_CONFIG_DIR="configs/c.${USB_CONFIG_INDEX}"
readonly USB_CONFIG_DIR

function usb_gadget_activate {
	ls /sys/class/udc >"${USB_DEVICE_PATH}/UDC"
	chmod 777 /dev/hidg0
}
export -f usb_gadget_activate

function usb_gadget_deactivate {
	echo '' >"${USB_DEVICE_PATH}/UDC"
}
export -f usb_gadget_deactivate
