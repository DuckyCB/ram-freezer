package main

import (
	"fmt"
	"os/exec"
)

func CopyRamImage() {
	fmt.Println("Copiando imagen de la ram...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/windows_copy_image")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}
