#!/usr/bin/env bash

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


printf "COMENZANDO INSTALACIÓN\n\n"

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

cd ./project-manager || { echo "Error: No se encontró el directorio project-manager"; exit 1; }
make build-project-manager
cd ..

cd ./ghost-keyboard/controller  || { echo "Error: No se encontró el directorio ghost-keyboard/controller"; exit 1; }
make build-ghost-keyboard
cd ../..

mkdir -p ./bin/scripts
cp ./ghost-keyboard/scripts/* ./bin/scripts


printf "CONFIGURANDO SISTEMA\n\n"

bash project-manager/setup.sh

bash ghost-keyboard/setup/setup.sh


printf "Reiniciando dispositivo en 10 segundos...\n"
sleep 10s
reboot
