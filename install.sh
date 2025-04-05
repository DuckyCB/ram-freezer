#!/usr/bin/env bash

# Echo commands to stdout.
#set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

clear
mkdir -p /opt/ram-freezer/bin/logs
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


log_info "Copiando scripts"
cd /opt/ram-freezer/

mkdir -p ./bin/scripts
cp ./ghost-keyboard/scripts/* ./bin/scripts


log_info "Configurando sistema..."
cd /opt/ram-freezer/

bash project-manager/setup/setup.sh
bash utils/usb-setup/setup.sh


log_info "Esperando 10 segundos antes de reiniciar. Crear /tmp/cancelar-reinicio para cancelar..."
echo "touch /tmp/cancelar-reinicio"

for i in {10..1}; do
    echo "Reinicio en $i segundos..."
    if [ -f /tmp/cancelar-reinicio ]; then
        log_info "Se detectó cancelación. Abortando reinicio."
        exit 0
    fi
    sleep 1
done

log_info "Reiniciando..."
sleep 1
reboot
