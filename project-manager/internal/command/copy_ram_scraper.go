package command

import (
	"os/exec"
	"project-manager/internal/logs"
)

func CopyRamScraper() {
	logs.Log.Info("Copiando ram-scraper...")

	cmd := exec.Command("/opt/ram-freezer/bin/ghost-keyboard", "-script", "/opt/ram-freezer/bin/scripts/windows_copy_ram_scraper")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(err.Error())
		return
	}

	logs.Log.Info(string(output))
}
