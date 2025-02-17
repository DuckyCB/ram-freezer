# Modulo RAM Scrapper - README

Este README proporciona instrucciones sobre cómo usar [WinPmem](https://github.com/Velocidex/WinPmem) para capturar la memoria de un sistema Windows.

## Descripción General

WinPmem es una herramienta ligera para capturar el contenido de la memoria RAM en sistemas Windows. Su objetivo es realizar una copia de la memoria volátil sin modificar el estado de la máquina.

### Componentes del Proyecto

- **winpmem_mini_x64_rc2.exe**: Herramienta para capturar la memoria del sistema.

## Requisitos

- Acceso a un sistema Windows con permisos de administrador.

## Pasos de Instalación y Configuración

1. **Descargar winpmem_mini_x64_rc2**
   - Descarga el archivo `winpmem_mini_x64_rc2.exe` desde el repositorio oficial o la página del proyecto.

2. **Transferir winpmem_mini_x64_rc2 a la máquina objetivo (Windows)**
   - Asegúrate de que tienes permisos de administrador en la máquina de destino.

## Uso de WinPmem para Capturar la Memoria

1. **Ejecutar winpmem_mini_x64_rc2**
   - Abre una consola de Windows (CMD) con permisos de administrador.
   - Navega a la ubicación de `winpmem_mini_x64_rc2.exe`.
   - Ejecuta el siguiente comando:

     ```cmd
     winpmem_mini_x64_rc2.exe RAM_winpmem.raw
     ```

     - Este proceso generará una copia completa de la memoria en el archivo `RAM_winpmem.raw`.

## Referencias

- [Documentación de WinPmem](https://github.com/Velocidex/WinPmem)
