#!/usr/bin/env bash

source /opt/ram-freezer/audit-trail/log.sh


log_info "Eliminando instalación"
bash /opt/ram-freezer/ghost-keyboard/setup/remove.sh
bash /opt/ram-freezer/project-manager/setup/remove.sh
