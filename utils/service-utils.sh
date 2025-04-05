#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh


function remove_service() {
  log_info "Eliminando $1"
  if systemctl is-active --quiet "$1"; then
    sudo systemctl stop "$1"
    if [ $? -eq 0 ]; then
      log_info "Servicio $1 detenido exitosamente."
      systemctl disable "$1"
      # si existe el servicio en /etc/systemd/system/ o /usr/lib/systemd/system/ lo eliminamos
      if [ -f /etc/systemd/system/"$1" ]; then
        log_info "Eliminando $1 de /etc/systemd/system/"
        rm /etc/systemd/system/"$1"
      fi
      if [ -f /usr/lib/systemd/system/"$1" ]; then
        log_info "Eliminando $1 de /usr/lib/systemd/system/"
        rm /usr/lib/systemd/system/"$1"
      fi
      systemctl daemon-reload
    else
      log_warn "No se pudo detener el servicio $1."
    fi
  elif systemctl is-enabled --quiet "$1"; then
    log_info "El servicio $1 está habilitado pero no activo."
    systemctl disable "$1"
    # si existe el servicio en /etc/systemd/system/ o /usr/lib/systemd/system/ lo eliminamos
    if [ -f /etc/systemd/system/"$1" ]; then
      log_info "Eliminando $1 de /etc/systemd/system/"
      rm /etc/systemd/system/"$1"
    fi
    if [ -f /usr/lib/systemd/system/"$1" ]; then
      log_info "Eliminando $1 de /usr/lib/systemd/system/"
      rm /usr/lib/systemd/system/"$1"
    fi
    systemctl daemon-reload
  else
    log_info "El servicio $1 no existe o no está habilitado."
  fi
}
