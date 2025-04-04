#!/usr/bin/env bash

# Echo commands to stdout.
#set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

clear
mkdir -p ./bin/logs
source /opt/ram-freezer/audit-trail/log.sh


if [[ "$UID" -ne 0 ]]; then
  log_fatal "Este script requiere privilegios de administrador (root)."
fi

log_info "Iniciando instalador de Ram Freezer"

printf "
     ██████╗  █████╗ ███╗   ███╗
     ██╔══██╗██╔══██╗████╗ ████║
     ██████╔╝███████║██╔████╔██║
     ██╔══██╗██╔══██║██║╚██╔╝██║
     ██║  ██║██║  ██║██║ ╚═╝ ██║
     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝

     ███████╗██████╗ ███████╗███████╗███████╗███████╗██████╗
     ██╔════╝██╔══██╗██╔════╝██╔════╝╚══███╔╝██╔════╝██╔══██╗
     █████╗  ██████╔╝█████╗  █████╗    ███╔╝ █████╗  ██████╔╝
     ██╔══╝  ██╔══██╗██╔══╝  ██╔══╝   ███╔╝  ██╔══╝  ██╔══██╗
     ██║     ██║  ██║███████╗███████╗███████╗███████╗██║  ██║
     ╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝
\n"

printf "by: fedeabdo & DuckyCB"

sleep 3
clear

# Permissions
log_info "Agregando permisos de ejecución"
chmod +x check.sh
chmod +x remove.sh

log_info "Instalando dependencias..."

# Verifica si Go está instalado
if ! command -v go &> /dev/null; then
  log_info "Instalando Go"
  sudo apt update && sudo apt install -y golang
fi

# Verificar si Make está instalado
if ! command -v make &> /dev/null; then
    log_info "Instalando Make"
    sudo apt update && sudo apt install -y make
fi

log_info "Compilando proyectos..."
cd /opt/ram-freezer/

cd ./project-manager || { log_error "No se encontró el directorio project-manager"; exit 1; }
log_info "Compilando Project Manager"
make build-project-manager
log_info "Project Manager fue compilado con exito"
cd ..

cd ./ghost-keyboard  || { log_error "No se encontró el directorio ghost-keyboard"; exit 1; }
log_info "Compilando Ghost Keyboard"
make build-ghost-keyboard
log_info "Ghost Keyboard fue compilado con exito"
cd ..

cd ./data-seal  || { log_error "No se encontró el directorio ghost-keyboard"; exit 1; }
log_info "Compilando Data Seal"
make build-data-seal
log_info "Data Seal fue compilado con exito"
cd ..

# Copy all scripts
log_info "Copiando scripts"
mkdir -p ./bin/scripts
cp ./ghost-keyboard/scripts/* ./bin/scripts


log_info "Configurando sistema..."
cd /opt/ram-freezer/

bash utils/usb-setup/setup.sh

log_info "Reiniciando dispositivo en 10 segundos... Presiona cualquier tecla para cancelar"
read -t 10 -n 1 -r -s key
if [ $? -eq 0 ]; then
    log_info "Reinicio cancelado. Es necesario reiniciar el sistema para que la instalación se complete"
else
    log_info "Reiniciando..."
    reboot
fi
