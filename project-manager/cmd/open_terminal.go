package main

import (
	"fmt"
	"os/exec"
)

func OpenTerminal() {
	fmt.Println("Abriendo terminal...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-f", "/opt/ram-freezer/bin/scripts/windows_open_terminal")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}