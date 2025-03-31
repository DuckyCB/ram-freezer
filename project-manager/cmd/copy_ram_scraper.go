package main

import (
	"fmt"
	"os/exec"
)

func CopyRamScraper() {
	fmt.Println("Copiando ram-scraper...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-f", "/opt/ram-freezer/bin/scripts/windows_copy_ram_scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}
