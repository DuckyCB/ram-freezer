#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh


if [[ "$UID" -ne 0 ]]; then
  log_install_fatal "Este script requiere privilegios de administrador (root)."
fi

log_install_info "Eliminando instalación"
bash /opt/ram-freezer/utils/usb-setup/remove.sh
bash /opt/ram-freezer/project-manager/setup/remove.sh
