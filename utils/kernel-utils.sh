#!/usr/bin/env bash

set -e
set -u

source /opt/ram-freezer/audit-trail/log.sh


check_modules() {
  libcomposite_loaded=true
  dwc2_loaded=true

  if ! lsmod | grep -q "libcomposite"; then
    log_error "El módulo libcomposite no está habilitado en el kernel."
    libcomposite_loaded=false
  fi

  if ! lsmod | grep -q "dwc2"; then
    log_error "El módulo dwc2 no está habilitado en el kernel."
    dwc2_loaded=false
  fi

  if $libcomposite_loaded && $dwc2_loaded; then
    log_info "Los módulos libcomposite y dwc2 están habilitados."
  fi
}

