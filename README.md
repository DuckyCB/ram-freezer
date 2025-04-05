
# RAM Freezer

## Setup

 - Instalar Raspberry Pi OS (Bookworm 32-bit)
   - [Raspberri Pi Imager](https://www.raspberrypi.com/software/) simplifica este proceso
   - Para la configuración inicial, se recomienda agregar la configuración de usuario, de red y habilitar la conexión SSH
 - (Recomendado) Conectar un almacenamiento USB a la Raspberry para copiar la imagen de la RAM
 - Conectarse a la raspberry
   - Por SSH es posible conectarse desde la misma red con `ssh usuario@raspberry.local`
 - Instalar sistema
```shell
curl -s -o- https://raw.githubusercontent.com/DuckyCB/ram-freezer/refs/heads/master/download.sh | sudo bash
```

