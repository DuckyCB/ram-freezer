#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh


check_modules() {
  if ! lsmod | grep -q "libcomposite"; then
    log_fatal "Error: El módulo libcomposite no está habilitado en el kernel."
    exit 1
  fi

  if ! lsmod | grep -q "dwc2"; then
    log_fatal "Error: El módulo dwc2 no está habilitado en el kernel."
    exit 1
  fi

  log_info "Los módulos libcomposite y dwc2 están habilitados."
}
