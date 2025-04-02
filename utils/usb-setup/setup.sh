#!/usr/bin/env bash

set -e
set -u

printf "\Configurando Gadget USB...\n"

# shellcheck source=/opt/ram-freezer/utils/usb-setup/usb-gadget.sh
source "/opt/ram-freezer/utils/usb-setup/usb-gadget.sh"


# Exec permissions
## USB
chmod +x "${USB_SETUP_PATH}/usb-modules-setup.sh"
chmod +x "${USB_SETUP_PATH}/init-usb-gadget.sh"
chmod +x "${USB_SETUP_PATH}/remove.sh"
chmod +x "${USB_SETUP_PATH}/remove-usb-gadget.sh"
## Keyboard
chmod +x "${KEYBOARD_SETUP_PATH}/init-keyboard.sh"
## Storage
chmod +x "${STORAGE_SETUP_PATH}/init-storage.sh"

# Load kernel modules
bash "${USB_SETUP_PATH}/usb-modules-setup.sh"

# Init USB device with modules
bash "${USB_SETUP_PATH}/init-usb-gadget.sh"

# Service
cp "${USB_SETUP_PATH}/usb-gadget.service" "${SYSTEM_SERVICES}/usb-gadget.service"
systemctl daemon-reload
systemctl enable usb-gadget.service


printf "Gadget USB configurado\n"
