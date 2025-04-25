package command

import (
	"fmt"
	"os/exec"
	"project-manager/internal/logs"
	"project-manager/pkg/utils"
	"time"
)

func CopyRamScraperToUSB() {
	logs.Log.Info("Copiando ram-scraper al USB...")

	// Desconecto el USB
	utils.DisconnectUSB()

	// Verificar si el USB está montado
	cmd := exec.Command("bash", "-c", "mount | grep '/dev/sda1'")
	output, err := cmd.CombinedOutput()

	if err != nil || len(output) == 0 {
		logs.Log.Info("El USB no está montado. Intentando montarlo...")
		utils.MountUSB()
	} else {
		logs.Log.Info("El USB ya está montado.")
	}

	// Crear el directorio en el USB
	logs.Log.Info("Creando el directorio en el USB...")
	cmd = exec.Command("sudo", "mkdir", "-p", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error creando el directorio: %s", string(output)))
		return
	}

	// Copiar los archivos de ram-scraper al USB
	logs.Log.Info("Copiando los archivos de ram-scraper al USB...")
	cmd = exec.Command("sudo", "cp", "-r", "/opt/ram-freezer/ram-scraper/.", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error copiando ram-scraper: %s", string(output)))
		return
	}
	// Esperar 3 segundos
	logs.Log.Info("Esperando 3 segundos...")
	time.Sleep(3 * time.Second)

	logs.Log.Info("Archivos copiados correctamente.")

	// Desmontar el USB
	utils.UmountUSB()

	utils.ConnectUSB()

	// montando el USB
	utils.MountUSB()

	logs.Log.Info("USB reconectado correctamente. Esperando 8 segundos")
	time.Sleep(8 * time.Second)

	logs.Log.Info("Proceso de copia de ram-scraper al USB completado.")
}
