# Ghost keyboard

## Requisitos

- Golang
    ```shell
    sudo apt install golang
    ```

## HID Setup

- Ghost keyboard se configura automaticamente al ejecutar el script de instalación principal de Ram Freezer

### Test setup

- Verificación de que los componentes existen y están cargados correctamente
    ```shell
    sudo go run /opt/ram-freezer/ghost-keyboard/setup/setup_check.go
    ```

- Verificación manual de que el teclado realmente escribe
    ```shell
    echo -ne "\x00\x00\x04\x00\x00\x00\x00\x00" > /dev/hidg0
    ```

## HID Controller

El sistema hecho con golang, utilizando los archivos `main.go` y `keycodes.go`, utiliza el dispositivo hid virtual para
escribir teclas emulando un teclado real.

### Build

Es posible compilar el sistema utilizando el comando `build` en el archivo Makefile. Este compila especificamente para
linux con arquitectura arm, utilizado por la raspberry.

Es posible compilar a mano utilizando el comando `build -o ghost-keyboard main.go keycodes.go`
> `ghost-keyboard` es el nombre del archivo final a ejecutar

### Run

El archivo compilado se guarda en la carpeta bin. Es posible ejecutarlo directamente con `./ghost-keyboard -script=file`
donde `file` es el archivo de input de texto a utilizar para escribir con el teclado virtual.

### Formato

El archivo `file` que será utilizado como input, admite tanto caracteres como teclas especiales.
Por ejemplo, si se quiere solo escribir caracteres, es posible hacerlo escribiendo algo tal que:

```
hola como estas?
```

Por otro lado, si se quiere escribir teclas especiales o combinaciones se puede hacer tal que:

```
{META + r}
```

> Esto presionara en simultaneo las teclas Meta (Super/Windows/Command) y la letra R

