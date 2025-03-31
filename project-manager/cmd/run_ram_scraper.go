package main

import (
	"fmt"
	"os/exec"
)

func RunRamScraper() {
	fmt.Println("Ejecutando ram-scraper...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-f", "/opt/ram-freezer/bin/scripts/windows_run_ram_scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}
