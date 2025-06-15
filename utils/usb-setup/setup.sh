#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh


log_install_info "Configurando Gadget USB..."

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

# remove old device
if [ -d "${USB_DEVICE_PATH}" ]; then
    log_install_info "Removing old gadget."
    bash "${USB_SETUP_PATH}/remove.sh"
fi

# Init USB device with modules
#bash "${USB_SETUP_PATH}/init-usb-gadget.sh"

# Service
cp "${USB_SETUP_PATH}/usb-gadget.service" "${SYSTEM_SERVICES}/usb-gadget.service"
systemctl daemon-reload
systemctl enable usb-gadget.service


log_install_info "Gadget USB configurado"
