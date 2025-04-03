#!/bin/bash

URL="https://github.com/DuckyCB/ram-freezer.git"
DIR="/opt/ram-freezer"

if [ -d "${DIR}" ]; then
  echo "El directorio ${DIR} existe. Actualizando..."
  cd "${DIR}"
  git pull origin master
else
  echo "El directorio ${DIR} no existe. Clonando..."
  sudo git clone "$URL" "$DIR"
fi

sudo chmod +x "$DIR/install.sh"
sudo bash "$DIR/install.sh"
