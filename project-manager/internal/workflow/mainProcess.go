package workflow

import (
	"os"
	"os/exec"
	"project-manager/internal/command"
	"project-manager/internal/files"
	"project-manager/internal/logs"
	"project-manager/pkg/utils"
	"ram-freezer/utils/rfutils/pkg/rfutils"
	"time"
)

// runSystem runs main system process
func (wfc *Controller) runSystem() {
	utils.MountUSB()

	command.CopyRamScraperToUSB()

	utils.UmountUSB()

	utils.ConnectUSB()
	command.OpenTerminal()

	logs.Log.Info("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	command.RunRamScraper()

	logs.Log.Info("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	utils.MountUSB()

	command.WaitAndValidateImage()

	utils.DisconnectUSB()
	utils.RemountUSB()

	dataSeal := "/opt/ram-freezer/bin/data-seal"
	logs.Log.Info("Comenzando creación de hashes de archivos")

	dataCmd := exec.Command(dataSeal, "-dir", "/mnt/usb/data/")
	dataCmd.Stdout = os.Stdout
	dataCmd.Stderr = os.Stderr
	if err := dataCmd.Run(); err != nil {
		logs.Log.Error(err.Error())
	}

	runPath := rfutils.GetOutPath()
	logsCmd := exec.Command(dataSeal, "-file", runPath+"/ram-freezer.log")
	logsCmd.Stdout = os.Stdout
	logsCmd.Stderr = os.Stderr
	if err := logsCmd.Run(); err != nil {
		logs.Log.Error(err.Error())
	}

	checksumCmd := exec.Command(dataSeal, "-checksum")
	checksumCmd.Stdout = os.Stdout
	checksumCmd.Stderr = os.Stderr
	if err := checksumCmd.Run(); err != nil {
		logs.Log.Error(err.Error())
	}

	files.CopyToUSB()

	// TODO hasta aca

	// TESSTING
	//command.TestKeyboard()
	//log.Println("espera 5 segundos para simular que esta haciendo algo")
	//time.Sleep(50 * time.Millisecond)
}
