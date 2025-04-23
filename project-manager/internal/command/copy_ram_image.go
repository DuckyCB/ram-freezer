package command

import (
	"os/exec"
	"project-manager/internal/logs"
)

func CopyRamImage() {
	logs.Log.Info("Copiando imagen de la ram...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/windows_copy_image")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(err.Error())
		return
	}

	logs.Log.Info(string(output))
}
