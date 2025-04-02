#!/usr/bin/env bash

set -e
set -u

readonly BASE_PATH='/opt/ram-freezer/utils/usb-setup'

chmod +x "$BASE_PATH/usb-modules-setup.sh"
chmod +x "$BASE_PATH/init-usb-gadget.sh"

"$BASE_PATH/usb-modules-setup.sh"

