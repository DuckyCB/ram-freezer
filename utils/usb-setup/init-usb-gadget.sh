#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh


#print_help() {
#  cat << EOF
#Usage: ${0##*/} [-h]
#Init USB gadget.
#  -h Display this help and exit.
#EOF
#}
#
## Parse command-line arguments.
#while getopts "h" opt; do
#  case "${opt}" in
#    h)
#      print_help
#      exit
#      ;;
#    *)
#      print_help >&2
#      exit 1
#  esac
#done

log_info "Iniciando usb-gadget"

# TODO: capaz esto puede ir en usb-modules-setup.sh ?
modprobe libcomposite

# USB device
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
echo "Ram Freezer" > "${USB_STRINGS_DIR}/product"

# Configs
mkdir -p "${USB_CONFIG_DIR}"
echo 250 > "${USB_CONFIG_DIR}/MaxPower"

CONFIGS_STRINGS_DIR="${USB_CONFIG_DIR}/${USB_STRINGS_DIR}"
mkdir -p "${CONFIGS_STRINGS_DIR}"
echo "Config ${USB_CONFIG_INDEX}: Keyboard and Storage" > "${CONFIGS_STRINGS_DIR}/configuration"

# Devices
log_info "Iniciando dispositivos de usb-gadget"
bash "${KEYBOARD_SETUP_PATH}/init-keyboard.sh"
bash "${STORAGE_SETUP_PATH}/init-storage.sh"

# Activate gadget
usb_gadget_activate

log_info "usb-gadget iniciado"
