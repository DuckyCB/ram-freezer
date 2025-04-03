#!/usr/bin/env bash

set -e
set -x
set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
readonly SCRIPT_DIR
# shellcheck source=utils/usb-gadget.sh
source "${SCRIPT_DIR}/utils/usb-gadget.sh"

# si existe el gadget USB, lo elimina
if [ -d "${USB_DEVICE_PATH}" ]; then
    echo "Gadget exists, removing it."
    chmod +x ${SCRIPT_DIR}/remove-usb-gadget.sh
    ${SCRIPT_DIR}/remove-usb-gadget.sh
fi

# se ejecuta el script de inicialización del gadget USB
chmod +x ${SCRIPT_DIR}/setup-usb-gadget.sh
${SCRIPT_DIR}/setup-usb-gadget.sh

# Se ejecuta el script de activación del gadget USB
usb_gadget_activate
