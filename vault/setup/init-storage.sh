#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh


log_info "Configurando Vault"

cd "${USB_DEVICE_PATH}"
mkdir -p "${USB_STORAGE_FUNCTIONS_DIR}"
echo 1 > "${USB_STORAGE_FUNCTIONS_DIR}/stall"
echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/removable"
## Use USB drive and create internal storage as fallback
if [ -e "${USB_DRIVE_PATH}" ]; then
    echo "${USB_DRIVE_PATH}" > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"
else
    log_warn "Almacenamiento USB no conectado, utilizando almacenamiento interno"

    if [ ! -e "${LOCAL_STORAGE_FILE}" ]; then
        dd if=/dev/zero of="${LOCAL_STORAGE_FILE}" bs=1M count="${LOCAL_STORAGE_SIZE}" status=progress
        mkdosfs "${LOCAL_STORAGE_FILE}" -F 32 -I
    else
        log_info "Usando archivo de almacenamiento existente: ${LOCAL_STORAGE_FILE}"
    fi

    echo "${LOCAL_STORAGE_FILE}" > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/file"
fi
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/cdrom"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/ro"
#echo 0 > "${USB_STORAGE_FUNCTIONS_DIR}/lun.0/nofua"


# Storage device label

# Ver si es FAT32 o exfat

if [ "$(blkid -o value -s TYPE "${USB_DRIVE_PATH}")" == "vfat" ]; then
    # Para FAT32
    echo "WARN : Formato de almacenamiento FAT32 detectado. En este formato los archivos no pueden ser mayores a 4GB."
    echo "WARN : Se recomienda usar exfat."
    echo "Renombrando dispositivo de almacenamiento a ${USB_STORAGE_DEVICE_NAME}"
    dosfslabel "${USB_DRIVE_PATH}" "${USB_STORAGE_DEVICE_NAME}"
elif [ "$(blkid -o value -s TYPE "${USB_DRIVE_PATH}")" == "exfat" ]; then
    # Para exfat
    echo "Renombrando dispositivo de almacenamiento a ${USB_STORAGE_DEVICE_NAME}"
    exfatlabel "${USB_DRIVE_PATH}" "${USB_STORAGE_DEVICE_NAME}"
else
    log_error "Formato de sistema de archivos no soportado. Solo FAT32 y exfat son soportados."
    exit 1
fi

log_info "Vault configurado"
