#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/service-utils.sh


log_install_info "Eliminando project-manager..."

# Service
remove_service "project-manager.service"

log_install_info "Project manager eliminado"
