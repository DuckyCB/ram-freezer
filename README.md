
# RAM Freezer

## Setup

### Requisitos previos

- Instalar Raspberry Pi OS (Bookworm 32-bit) en la Raspberry

> [Raspberri Pi Imager](https://www.raspberrypi.com/software/) simplifica este proceso

> Para la configuración inicial, se recomienda agregar la configuración de usuario, de red y habilitar la conexión SSH

- Conectar un LED en el GPIO 27 (utilizar resistencia de 220Ω)

- Conectar un botón en el GPIO 17 de la Raspberry

- (Recomendado) Conectar un almacenamiento USB a la Raspberry

### Conectarse a la raspberry

#### Fisicamente

- Utilizando un monitor, teclado y mouse conectados directamente a la raspberry

#### Red

- Conectarse a la misma red que la raspberry

- Ejecutar `ssh usuario@raspberry.local` (donde usuario es el nombre de usuario asignado, y raspberry es el nombre del host)

> Esto permite conectarse por `ssh` y controlar la terminal

### Instalar y actualizar sistema

> Desde la raspberry (ver Conectarse a la Raspberry)

```shell
curl -s -o- https://raw.githubusercontent.com/DuckyCB/ram-freezer/refs/heads/master/download.sh | sudo bash
```

## Como se usa

### Estado del LED

- Parpadeo lento (cada 1 segundo): El sistema se encuentra activo esperando a que se pulse el botón

- Parpadeo medio (cada 0.5 segundos): El sistema se está ejecutando, esperar a que termine

- Parpadeo rápido (cada 0.1 segundos, durante 10 segundos): El sistema completó la ejecución. Volverá al estado inicial esperando a que se pulse el botón.

### Iniciar ejecución

- Conectar a la PC objetivo 

- Presionar el botón

> La ejecución puede demorar varios minutos, esta depende de la cantidad de RAM que tiene la PC
