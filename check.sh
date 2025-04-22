#!/usr/bin/env bash


source /opt/ram-freezer/audit-trail/log.sh


if [[ "$UID" -ne 0 ]]; then
  log_fatal "Este script requiere privilegios de administrador (root)."
fi

log_info "Iniciando verificación de la instalación"

go run "/opt/ram-freezer/project-manager/setup/setup_check.go"

log_info "Ghost keyboard"
go run "/opt/ram-freezer/ghost-keyboard/setup/setup_check.go"

log_info "Vault"
go run "/opt/ram-freezer/vault/setup/setup_check.go"
