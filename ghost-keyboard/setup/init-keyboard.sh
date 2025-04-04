#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh


log_info "Configurando ghost keyboard"

mkdir -p "${USB_KEYBOARD_FUNCTIONS_DIR}"
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

ln -s "${USB_KEYBOARD_FUNCTIONS_DIR}" "${USB_CONFIG_DIR}/"

log_info "Ghost Keyboard configurado"
