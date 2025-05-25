package workflow

import (
	"project-manager/internal/command"
	"project-manager/internal/logs"
	"project-manager/pkg/utils"
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

	command.HashFiles("/mnt/usb/data/")

	// TODO hasta aca

	// TESSTING
	//command.TestKeyboard()
	//log.Println("espera 5 segundos para simular que esta haciendo algo")
	//time.Sleep(50 * time.Millisecond)
}
