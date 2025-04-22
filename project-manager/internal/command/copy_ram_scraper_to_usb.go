package command

import (
	"fmt"
	"os/exec"
	"time"
	"project-manager/pkg/utils"
)

func CopyRamScraperToUSB() {
	fmt.Println("Copiando ram-scraper al USB...")

	// Desconecto el USB
	utils.DisconnectUSB()
	// fmt.Println("Reconectando el USB...")
	// cmd := exec.Command("bash", "-c", "echo '' | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error reconectando el USB:", string(output))
	// 	return
	// }


	// Verificar si el USB está montado
	cmd := exec.Command("bash", "-c", "mount | grep '/dev/sda1'")
	output, err := cmd.CombinedOutput()

	if err != nil || len(output) == 0 {
		fmt.Println("El USB no está montado. Intentando montarlo...")

		// Montar el USB
		utils.MountUSB()

	} else {
		fmt.Println("El USB ya está montado.")
	}

	// Crear el directorio en el USB
	fmt.Println("Creando el directorio en el USB...")
	cmd = exec.Command("sudo", "mkdir", "-p", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error creando el directorio:", string(output))
		return
	}

	// Copiar los archivos de ram-scraper al USB
	fmt.Println("Copiando los archivos de ram-scraper al USB...")
	cmd = exec.Command("sudo", "cp", "-r", "/opt/ram-freezer/ram-scraper/.", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error copiando ram-scraper:", string(output))
		return
	}
	// Esperar 3 segundos
	fmt.Println("Esperando 3 segundos...")
	time.Sleep(3 * time.Second)

	fmt.Println("Archivos copiados correctamente.")

	// Desmontar el USB
	utils.UmountUSB()
	
	utils.ConnectUSB()

	// montando el USB
	utils.MountUSB()

	
	fmt.Println("USB reconectado correctamente. Esperando 8 segundos")
	time.Sleep(8 * time.Second)

	fmt.Println("Proceso de copia de ram-scraper al USB completado.")
}
