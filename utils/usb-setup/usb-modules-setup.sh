#!/usr/bin/env bash

set -e
set -u

source "/opt/ram-freezer/utils/usb-setup/usb-gadget.sh"


if ! grep --quiet '^dtoverlay=dwc2$' "${CONFIG_FILE}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${CONFIG_FILE}"
fi

if ! grep --quiet '^dwc2$' "${MODULES_PATH}" ; then
  echo 'dwc2' | tee --append "${MODULES_PATH}"
fi

if ! grep --quiet '^dtoverlay=dwc2$' "${CONFIG_FILE}" ; then
  echo 'dtoverlay=dwc2' | tee --append "${CONFIG_FILE}"
fi
