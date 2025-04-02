package main

import (
	"fmt"
	"os/exec"
)

func CopyRamScraperToUSB() {
	fmt.Println("Copiando ram-scraper al USB...")

	// Montar el USB si no esta montado
	cmd := exec.Command("lsblk", "-o", "NAME,MOUNTPOINT")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error al listar los dispositivos:", err)
		return
	}
	fmt.Println(string(output))

	// Verificar si el USB está montado
	if string(output) == "" {
		fmt.Println("El USB no está montado. Montando el USB...")
			

		fmt.Println("Montando el USB...")
		cmd := exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error montando el USB:", err)
			return
		}
		fmt.Println(string(output))
	} else {
		fmt.Println("El USB ya está montado.")
	}

	// Crear el directorio en el USB
	fmt.Println("Creando el directorio en el USB...")
	cmd = exec.Command("sudo", "mkdir", "-p", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error creando el directorio:", err)
		return
	}

	// Copiar los archivos de ram-scraper al USB
	fmt.Println("Copiando los archivos de ram-scraper al USB...")
	cmd = exec.Command("sudo", "cp", "-r", "/opt/ram-freezer/", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error copiando ram-scraper:", err)
		return
	}
	fmt.Println(string(output))


	// Desmontar el USB
	fmt.Println("Desmontando el USB...")
	cmd = exec.Command("sudo", "umount", "/mnt/usb/")

	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error desmontando el USB:", err)
		return
	}
	fmt.Println(string(output))

	// Reconectando el USB
	fmt.Println("Reconectando el USB...")
	// echo fe980000.usb | sudo tee UDC
	cmd = exec.Command("echo", "fe980000.usb", "|", "sudo", "tee", "/sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error reconectando el USB:", err)
		return
	}
	fmt.Println(string(output))

	// wait 1s
	cmd = exec.Command("sleep", "1")

	// ls /sys/class/udc
	cmd = exec.Command("ls", "/sys/class/udc", "|", "sudo", "tee", "/sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error listando el USB:", err)
		return
	}
	fmt.Println(string(output))

	// sleep 5
	cmd = exec.Command("sleep", "5")

	fmt.Println("Proceso de copia de ram-scraper al USB completado.")
}
