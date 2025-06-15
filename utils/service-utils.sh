#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/file-utils.sh


function remove_service() {
  log_install_info "Eliminando $1"
  if systemctl is-active --quiet "$1"; then
    sudo systemctl stop "$1"
    if [ $? -eq 0 ]; then
      log_install_info "Servicio $1 detenido exitosamente."
      systemctl disable "$1"
      remove_file /etc/systemd/system/"$1"
      remove_file /usr/lib/systemd/system/"$1"
      systemctl daemon-reload
    else
      log_install_warn "No se pudo detener el servicio $1."
    fi
  elif systemctl is-enabled --quiet "$1"; then
    log_install_info "El servicio $1 está habilitado pero no activo."
    systemctl disable "$1"
    remove_file /etc/systemd/system/"$1"
    remove_file /usr/lib/systemd/system/"$1"
    systemctl daemon-reload
  else
    log_install_info "El servicio $1 no existe o no está habilitado."
  fi
}
