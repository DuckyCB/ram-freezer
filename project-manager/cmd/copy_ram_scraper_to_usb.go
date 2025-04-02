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
}