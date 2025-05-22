#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh
source /opt/ram-freezer/audit-trail/log.sh


log_install_info "Configurando modulos del Kernel"

log_install_info "Cargando dwc2"
if ! grep --quiet '^dtoverlay=dwc2$' "${CONFIG_FILE}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${CONFIG_FILE}"
fi
if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi
log_install_info "dwc2 cargado en el kernel"
