package command

import (
	"fmt"
	"os/exec"
	"time"
)

func CopyRamScraperToUSB() {
	fmt.Println("Copiando ram-scraper al USB...")

	// Desconecto el USB
	fmt.Println("Reconectando el USB...")
	cmd := exec.Command("bash", "-c", "echo '' | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error reconectando el USB:", string(output))
		return
	}


	// Verificar si el USB está montado
	cmd = exec.Command("bash", "-c", "mount | grep '/dev/sda1'")
	output, err = cmd.CombinedOutput()

	if err != nil || len(output) == 0 {
		fmt.Println("El USB no está montado. Intentando montarlo...")

		cmd = exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
		output, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error montando el USB:", string(output))
			return
		}
		fmt.Println("USB montado correctamente.")
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
	fmt.Println("Desmontando el USB...")
	cmd = exec.Command("sudo", "umount", "/mnt/usb/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error desmontando el USB:", string(output))
		return
	}
	fmt.Println("USB desmontado correctamente.")

	cmd = exec.Command("bash", "-c", "ls /sys/class/udc | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error listando el USB:", string(output))
		return
	}

	// montando el USB
	fmt.Println("Montando el USB...")
	cmd = exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error montando el USB:", string(output))
		return
	}
	fmt.Println("USB montado correctamente.")

	fmt.Println("USB reconectado correctamente. Esperando 8 segundos")
	time.Sleep(8 * time.Second)

	fmt.Println("Proceso de copia de ram-scraper al USB completado.")
}
