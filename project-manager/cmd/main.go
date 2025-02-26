package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	log.Println("Starting project manager")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "help")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))

}
