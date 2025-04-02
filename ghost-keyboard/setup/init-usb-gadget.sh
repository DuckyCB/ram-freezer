#!/usr/bin/env bash

set -e 
set -x
set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
readonly SCRIPT_DIR
# shellcheck source=utils/usb-gadget.sh
source "${SCRIPT_DIR}/utils/usb-gadget.sh"

chmod +x setup-usb-gadget.sh
./setup-usb-gadget.sh

# print pwd
echo "Current directory: $(pwd)"

usb_gadget_activate
