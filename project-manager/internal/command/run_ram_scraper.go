package command

import (
	"fmt"
	"os/exec"
	"project-manager/internal/logs"
)

func RunRamScraper() {
	logs.Log.Info("Ejecutando ram-scraper...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/windows_run_ram_scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(err.Error())
		return
	}

	logs.Log.Info(string(output))
}

func WaitAndValidateImage() {
	logs.Log.Info("Esperando la creacion y validacion de la imagen RAM...")

	cmd := exec.Command("/opt/ram-freezer/bin/ram-scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error ejecutando el binario: %v", err))
		return
	}

	logs.Log.Info(string(output))
}
