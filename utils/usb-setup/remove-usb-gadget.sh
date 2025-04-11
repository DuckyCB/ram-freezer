#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh

IS_MOUNTED=$(mount | grep "${USB_DEVICE_DIR}" | awk '{print $3}')
if [ -n "${IS_MOUNTED}" ]; then
    log_info "Desmontando USB ${USB_DRIVE_PATH} de ${USB_MOUNT_POINT}"
    umount "${USB_MOUNT_POINT}"
fi

log_info "Eliminando usb-gadget"

cd "${USB_GADGET_PATH}"

if [ ! -d "${USB_DEVICE_DIR}" ]; then
    log_info "Gadget does not exist, quitting."
    exit 0
fi

pushd "${USB_DEVICE_DIR}"

# Disable all gadgets
if [ -n "$(cat UDC)" ]; then
  echo "" > UDC
fi

# Walk items in `configs`
for config in ${USB_ALL_CONFIGS_DIR} ; do
    # Exit early if there are no entries
    [ "${config}" == "${USB_ALL_CONFIGS_DIR}" ] && break

    # Remove all functions from config
    for function in ${USB_ALL_FUNCTIONS_DIR} ; do
      file="${config}/$(basename "${function}")"
      [ -e "${file}" ] && rm "${file}"
    done

    # Remove strings in config
    [ -d "${config}/${USB_STRINGS_DIR}" ] && rmdir "${config}/${USB_STRINGS_DIR}"

    rmdir "${config}"
done

# Remove functions
for function in ${USB_ALL_FUNCTIONS_DIR} ; do
    [ -d "${function}" ] && rmdir "${function}"
done

# Remove strings from device
[ -d "${USB_STRINGS_DIR}" ] && rmdir "${USB_STRINGS_DIR}"

popd

rmdir "${USB_DEVICE_DIR}"

log_info "usb-gadget eliminado"
