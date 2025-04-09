package command

import (
	"fmt"
	"os/exec"
)

func ValidateRamImage() {
	fmt.Println("Ejecutando ram-scraper-validator...")

	// TODO: el pendrive paso a llamarse USB_VAULT, si esto se rompe es por eso
	cmd := exec.Command("/opt/ram-freezer/bin/ram-scraper-validator")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}
