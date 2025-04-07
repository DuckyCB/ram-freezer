package command

import (
	"fmt"
	"os/exec"
)

func TestKeyboard() {
	fmt.Println("Probando ghost keyboard...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/test")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el binario:", err)
		return
	}

	fmt.Println(string(output))
}
