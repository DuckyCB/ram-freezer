#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh


log_install_info "Iniciando configuración de Project Manager"

cd /opt/ram-freezer/project-manager/setup


log_install_info "Creando servicio project-manager.service"

# TODO: si existe un servicio viejo, lo deberia eliminar

cp project-manager.service /lib/systemd/system/project-manager.service
systemctl daemon-reload
systemctl enable project-manager.service


chmod +x remove.sh

log_install_info "Project Manager configurado correctamente"
