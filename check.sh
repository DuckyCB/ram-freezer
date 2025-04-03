#!/usr/bin/env bash


source /home/ducky/Projects/tesis/ram-freezer/audit-trail/log.sh

log_info "Iniciando verificación de la instalación\n\n"

log_info "Ghost keyboard:\n"
go run /opt/ram-freezer/ghost-keyboard/setup/setup_check.go
