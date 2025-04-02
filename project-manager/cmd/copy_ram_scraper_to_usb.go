package main

import (
	"fmt"
	"os/exec"
)

func CopyRamScraperToUSB() {
	fmt.Println("Copiando ram-scraper al USB...")

	// Montar el USB
	fmt.Println("Montando el USB...")
	cmd := exec.Command("mount", "/dev/sda1", "/mnt/usb/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error montando el USB:", err)
		return
	}
	fmt.Println(string(output))

	// Crear el directorio en el USB
	fmt.Println("Creando el directorio en el USB...")
	cmd = exec.Command("mkdir", "-p", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error creando el directorio:", err)
		return
	}

	// Copiar los archivos de ram-scraper al USB
	fmt.Println("Copiando los archivos de ram-scraper al USB...")
	cmd = exec.Command("cp", "-r", "/opt/ram-freezer/", "/mnt/usb/ram-scraper/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error copiando ram-scraper:", err)
		return
	}
	fmt.Println(string(output))


	// Desmontar el USB
	fmt.Println("Desmontando el USB...")
	cmd = exec.Command("umount", "/mnt/usb/")

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
	cmd = exec.Command("ls", "/sys/class/udc", "tee", "/sys/kernel/config/usb_gadget/ram-freezer/UDC")
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
