package workflow

import (
	"project-manager/internal/command"
	"project-manager/internal/logs"
	"project-manager/pkg/utils"
	"time"
)

// runSystem runs main system process
func (wfc *Controller) runSystem() {
	// Llama a las funciones en el orden deseado

	utils.MountUSB()

	// Copiar archivos de ram-scraper al USB
	command.CopyRamScraperToUSB()

	utils.UmountUSB()

	utils.ConnectUSB()
	// Abrir la terminal
	command.OpenTerminal()
	
	logs.Log.Info("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	// Crear la imagen de RAM
	command.RunRamScraper()

	logs.Log.Info("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	utils.MountUSB()
	
	// Validar la imagen de RAM
	command.WaitAndValidateImage()

	// Reconecto el USB
	utils.DisconnectUSB()
	utils.RemountUSB()

	// Crear el hash de la imagen de RAM
	command.HashFiles("/mnt/usb/data/")

	// TODO hasta aca

	// TESSTING
	//command.TestKeyboard()
	//log.Println("espera 5 segundos para simular que esta haciendo algo")
	//time.Sleep(50 * time.Millisecond)
}
