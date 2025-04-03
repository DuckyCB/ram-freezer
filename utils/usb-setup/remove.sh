#!/usr/bin/env bash

set -e
set -u


printf "Eliminando usb-gadget...\n"

source "/opt/ram-freezer/utils/usb-setup/usb-gadget.sh"


# Service
# TODO convertir esto en una funcion reutilizable
if systemctl is-active --quiet usb-gadget.service; then
  sudo systemctl stop usb-gadget.service
  if [ $? -eq 0 ]; then
    echo "Servicio usb-gadget.service detenido exitosamente."
    systemctl disable usb-gadget.service
    rm /etc/systemd/system/usb-gadget.service
    rm /usr/lib/systemd/system/usb-gadget.service
    systemctl daemon-reload
  else
    echo "Error al detener el servicio usb-gadget.service."
  fi
elif systemctl is-enabled --quiet usb-gadget.service; then
  echo "El servicio usb-gadget.service está habilitado pero no activo."
  systemctl disable usb-gadget.service
  rm /etc/systemd/system/usb-gadget.service
  rm /usr/lib/systemd/system/usb-gadget.service
  systemctl daemon-reload
else
  echo "El servicio usb-gadget.service no existe o no está habilitado."
fi

# Gadget
bash /opt/ram-freezer/utils/usb-setup/remove-usb-gadget.sh


printf "usb-gadget eliminado\n"
