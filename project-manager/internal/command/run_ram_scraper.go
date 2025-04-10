package command

import (
	"fmt"
	"os/exec"
)

func RunRamScraper() {
	fmt.Println("Ejecutando ram-scraper...")

	// TODO: el pendrive paso a llamarse USB_VAULT, si esto se rompe es por eso
	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/windows_run_ram_scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}

func WaitAndValidateImage() {
	fmt.Println("Esperando la creacion y validacion de la imagen RAM...")

	cmd := exec.Command("/opt/ram-freezer/bin/ram-scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}