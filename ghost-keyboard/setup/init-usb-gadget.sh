#!/usr/bin/env bash

set -e
set -x
set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
readonly SCRIPT_DIR
# shellcheck source=utils/usb-gadget.sh
source "${SCRIPT_DIR}/utils/usb-gadget.sh"

print_help() {
  cat << EOF
Usage: ${0##*/} [-h]
Init USB gadget.
  -h Display this help and exit.
EOF
}

# Parse command-line arguments.
while getopts "h" opt; do
  case "${opt}" in
    h)
      print_help
      exit
      ;;
    *)
      print_help >&2
      exit 1
  esac
done

modprobe libcomposite

cd "${USB_GADGET_PATH}"
mkdir -p "${USB_DEVICE_DIR}"
cd "${USB_DEVICE_DIR}"

echo 0x1d6b > idVendor  # Linux Foundation
echo 0x0104 > idProduct # Multifunction Composite Gadget
echo 0x0100 > bcdDevice # v1.0.0
echo 0x0200 > bcdUSB    # USB2

mkdir -p "$USB_STRINGS_DIR"
echo "22slun7emp6l8qzrocc4" > "${USB_STRINGS_DIR}/serialnumber"
echo "Ram Freezer" > "${USB_STRINGS_DIR}/manufacturer"
echo "Ghost Keyboard" > "${USB_STRINGS_DIR}/product"

# Keyboard
mkdir -p "$USB_KEYBOARD_FUNCTIONS_DIR"
echo 1 > "${USB_KEYBOARD_FUNCTIONS_DIR}/protocol" # Keyboard
echo 1 > "${USB_KEYBOARD_FUNCTIONS_DIR}/subclass" # Boot interface subclass
echo 8 > "${USB_KEYBOARD_FUNCTIONS_DIR}/report_length"

# Write the report descriptor
D=$(mktemp)
{
  echo -ne \\x05\\x01       # Usage Page (Generic Desktop Ctrls)
  echo -ne \\x09\\x06       # Usage (Keyboard)
  echo -ne \\xA1\\x01       # Collection (Application)
  echo -ne \\x05\\x08       #   Usage Page (LEDs)
  echo -ne \\x19\\x01       #   Usage Minimum (Num Lock)
  echo -ne \\x29\\x03       #   Usage Maximum (Scroll Lock)
  echo -ne \\x15\\x00       #   Logical Minimum (0)
  echo -ne \\x25\\x01       #   Logical Maximum (1)
  echo -ne \\x75\\x01       #   Report Size (1)
  echo -ne \\x95\\x03       #   Report Count (3)
  echo -ne \\x91\\x02       #   Output (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position,Non-volatile)
  echo -ne \\x09\\x4B       #   Usage (Generic Indicator)
  echo -ne \\x95\\x01       #   Report Count (1)
  echo -ne \\x91\\x02       #   Output (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position,Non-volatile)
  echo -ne \\x95\\x04       #   Report Count (4)
  echo -ne \\x91\\x01       #   Output (Const,Array,Abs,No Wrap,Linear,Preferred State,No Null Position,Non-volatile)
  echo -ne \\x05\\x07       #   Usage Page (Kbrd/Keypad)
  echo -ne \\x19\\xE0       #   Usage Minimum (0xE0)
  echo -ne \\x29\\xE7       #   Usage Maximum (0xE7)
  echo -ne \\x95\\x08       #   Report Count (8)
  echo -ne \\x81\\x02       #   Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
  echo -ne \\x75\\x08       #   Report Size (8)
  echo -ne \\x95\\x01       #   Report Count (1)
  echo -ne \\x81\\x01       #   Input (Const,Array,Abs,No Wrap,Linear,Preferred State,No Null Position)
  echo -ne \\x19\\x00       #   Usage Minimum (0x00)
  echo -ne \\x29\\x91       #   Usage Maximum (0x91)
  echo -ne \\x26\\xFF\\x00  #   Logical Maximum (255)
  echo -ne \\x95\\x06       #   Report Count (6)
  echo -ne \\x81\\x00       #   Input (Data,Array,Abs,No Wrap,Linear,Preferred State,No Null Position)
  echo -ne \\xC0            # End Collection
} >> "$D"
cp "$D" "${USB_KEYBOARD_FUNCTIONS_DIR}/report_desc"
# Enable pre-boot events (if the gadget driver supports it).
if [[ -f "${USB_KEYBOARD_FUNCTIONS_DIR}/no_out_endpoint" ]]; then
  echo 1 > "${USB_KEYBOARD_FUNCTIONS_DIR}/no_out_endpoint"
fi

# Storage
mkdir -p "${USB_STORAGE_FUNCTIONS_DIR}"
echo 1 > "${USB_STORAGE_FUNCTIONS_DIR}/stall"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/removable"
echo /dev/sda1 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"


mkdir -p "${USB_CONFIG_DIR}"
echo 250 > "${USB_CONFIG_DIR}/MaxPower"

CONFIGS_STRINGS_DIR="${USB_CONFIG_DIR}/${USB_STRINGS_DIR}"
mkdir -p "${CONFIGS_STRINGS_DIR}"
echo "Config ${USB_CONFIG_INDEX}: ECM network" > "${CONFIGS_STRINGS_DIR}/configuration"

ln -s "${USB_STORAGE_FUNCTIONS_DIR}" "${USB_CONFIG_DIR}/"
ln -s "${USB_KEYBOARD_FUNCTIONS_DIR}" "${USB_CONFIG_DIR}/"

usb_gadget_activate

