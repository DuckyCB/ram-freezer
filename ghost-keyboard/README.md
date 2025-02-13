
# Ghost keyboard

## Requisitos

- Golang

```shell
sudo apt install golang
```

## HID Setup

- Darle permisos de ejecución a `setup.sh` con `sudo chown +x setup.sh`
- Ejecutar `setup.sh` con `./setup.sh`
- Reiniciar dispositivo con `sudo reboot`

### Test setup

```shell
go run setup_check.go
```

## HID Controller

El sistema hecho con golang, utilizando los archivos `main.go` y `keycodes.go`, utiliza el dispositivo hid virtual para escribir teclas emulando un teclado real.

### Build

Es posible compilar el sistema utilizando el comando `build` en el archivo Makefile. Este compila especificamente para linux con arquitectura arm, utilizado por la raspberry.

Es posible compilar a mano utilizando el comando `build -o ghost-keyboard main.go keycodes.go`
> `ghost-keyboard` es el nombre del archivo final a ejecutar

### Run

El archivo compilado se guarda en la carpeta bin. Es posible ejecutarlo directamente con `./ghost-keyboard -f file` donde `file` es el archivo de input de texto a utilizar para escribir con el teclado virtual.

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

