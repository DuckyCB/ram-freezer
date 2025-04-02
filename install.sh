#!/usr/bin/env bash

# Echo commands to stdout.
#set -x
# Exit on first error.
set -e
# Treat undefined environment variables as errors.
set -u

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

printf "by: fedeabdo & DuckyCB\n\n\n"

printf "Comenzando instalación\n\n"

printf "INSTALANDO DEPENDENCIAS\n\n"

# Verifica si Go está instalado
if ! command -v go &> /dev/null; then
  printf "Instalando Go...\n"
  sudo apt update && sudo apt install -y golang
fi

# Verificar si Make está instalado
if ! command -v make &> /dev/null; then
    printf "Instalando Make...\n"
    sudo apt update && sudo apt install -y make
fi

printf "COMPILANDO PROYECTOS\n\n"
cd /opt/ram-freezer/

cd ./project-manager || { echo "Error: No se encontró el directorio project-manager"; exit 1; }
make build-project-manager
cd ..

cd ./ghost-keyboard  || { echo "Error: No se encontró el directorio ghost-keyboard"; exit 1; }
make build-ghost-keyboard
cd ..

cd ./data-seal  || { echo "Error: No se encontró el directorio ghost-keyboard"; exit 1; }
make build-data-seal
cd ..

# Copy all scripts
mkdir -p ./bin/scripts
cp ./ghost-keyboard/scripts/* ./bin/scripts


printf "CONFIGURANDO SISTEMA\n\n"
cd /opt/ram-freezer/

bash utils/usb-setup/setup.sh

# Permissions
chmod +x check.sh
chmod +x remove.sh

printf "Reiniciando dispositivo en 10 segundos...\nPresiona cualquier tecla para cancelar\n"
if read -t 10 -n 1; then
    printf "Reinicio cancelado\n> Es necesario reiniciar el sistema para que la instalación se complete"
else
    printf "Reiniciando...\n"
    reboot
fi
