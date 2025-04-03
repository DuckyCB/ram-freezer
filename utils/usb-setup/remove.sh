#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh
source /opt/ram-freezer/utils/usb-setup/usb-gadget.sh
source /opt/ram-freezer/utils/service-utils.sh


log_info "Eliminando usb-gadget..."

# Service
remove_service "usb-gadget.service"

# Gadget
bash /opt/ram-freezer/utils/usb-setup/remove-usb-gadget.sh

log_info "usb-gadget eliminado correctamente"
