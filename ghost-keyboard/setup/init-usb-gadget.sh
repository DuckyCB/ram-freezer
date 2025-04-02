#!/usr/bin/env bash

set -e
set -x
set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
readonly SCRIPT_DIR
# shellcheck source=utils/usb-gadget.sh
source "${SCRIPT_DIR}/utils/usb-gadget.sh"

chmod +x ${SCRIPT_DIR}/setup-usb-gadget.sh
${SCRIPT_DIR}/setup-usb-gadget.sh

# Se ejecuta el script de activación del gadget USB
usb_gadget_activate
