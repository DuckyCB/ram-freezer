#!/usr/bin/env bash

# Echo commands to stdout.
#set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

clear

source /opt/ram-freezer/audit-trail/log.sh

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

printf "by: fedeabdo & DuckyCB\n\n"

sleep 2
clear

log_info "Instalando dependencias...\n\n"

# Verifica si Go está instalado
if ! command -v go &> /dev/null; then
  log_info "Instalando Go\n"
  sudo apt update && sudo apt install -y golang
fi

# Verificar si Make está instalado
if ! command -v make &> /dev/null; then
    log_info "Instalando Make\n"
    sudo apt update && sudo apt install -y make
fi

log_info "Compilando proyectos...\n\n"
cd /opt/ram-freezer/

cd ./project-manager || { log_error "Error: No se encontró el directorio project-manager"; exit 1; }
log_info "Compilando Project Manager"
make build-project-manager
log_info "Project Manager fue compilado con exito"
cd ..

cd ./ghost-keyboard  || { log_error "Error: No se encontró el directorio ghost-keyboard"; exit 1; }
log_info "Compilando Ghost Keyboard"
make build-ghost-keyboard
log_info "Ghost Keyboard fue compilado con exito"
cd ..

cd ./data-seal  || { log_error "Error: No se encontró el directorio ghost-keyboard"; exit 1; }
log_info "Compilando Data Seal"
make build-data-seal
log_info "Data Seal fue compilado con exito"
cd ..

# Copy all scripts
log_info "Copiando scripts"
mkdir -p ./bin/scripts
cp ./ghost-keyboard/scripts/* ./bin/scripts


log_info "Configurando sistema...\n\n"
cd /opt/ram-freezer/

bash utils/usb-setup/setup.sh

# Permissions
log_info "Agregando permisos de ejecución\n\n"
chmod +x check.sh
chmod +x remove.sh

log_info "Reiniciando dispositivo en 10 segundos...\nPresiona cualquier tecla para cancelar\n"
if read -t 10 -n 1; then
    log_info "Reinicio cancelado\n Es necesario reiniciar el sistema para que la instalación se complete"
else
    log_info "Reiniciando...\n"
    reboot
fi
