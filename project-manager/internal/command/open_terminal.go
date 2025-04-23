package command

import (
	"fmt"
	"os/exec"
	"project-manager/internal/logs"
)

func OpenTerminal() {
	logs.Log.Info("Abriendo terminal...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/windows_open_terminal")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error ejecutando el binario: %v", err))
		return
	}

	logs.Log.Info(string(output))
}
