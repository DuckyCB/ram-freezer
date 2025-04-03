#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh


log_info "Iniciando configuración de Project Manager"

cd /opt/ram-freezer/project-manager/setup

# Start project-manager at boot
log_info "Creando servicio project-manager.service"
cp project-manager.service /lib/systemd/system/project-manager.service
systemctl daemon-reload
systemctl enable project-manager.service

chmod +x remove.sh

log_info "Project Manager configurado correctamente"
