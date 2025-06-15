#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh


function remove_file() {
  if [ -f "$1" ]; then
    rm "$1"
  else
    log_install_info "El archivo '$1' no existe. No se puede borrar."
  fi
}