package main

import (
	"fmt"
)

func CopyRamScraperToUSB() {
	fmt.Println("Copiando ram-scraper al USB...")

	// Montar el USB
	cmd := exec.Command("mount", "/dev/sda1", "/mnt/usb/")

	// Crear el directorio en el USB
	cmd := exec.Command("mkdir", "-p", "/mnt/usb/ram-scraper/")

	cmd := exec.Command("cp", "-r", "/opt/ram-freezer/", "/mnt/usb/ram-scraper/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error copiando ram-scraper:", err)
		return
	}
	fmt.Println(string(output))
}