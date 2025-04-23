package command

import (
	"fmt"
	"os/exec"
	"project-manager/internal/logs"
)

func TestKeyboard() {
	logs.Log.Info("Probando ghost keyboard...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/test")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error ejecutando el binario: %v", err))
		return
	}

	logs.Log.Info(string(output))
}
