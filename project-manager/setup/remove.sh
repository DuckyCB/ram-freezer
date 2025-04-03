#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/service-utils.sh


log_info "Eliminando project-manager...\n"

# Service
remove_service "project-manager.service"

log_info "Project manager eliminado\n"
