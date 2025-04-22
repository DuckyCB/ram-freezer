package workflow

import (
	"log"
	"project-manager/internal/command"
	"project-manager/pkg/utils"
	"time"
)

// runSystem runs main system process
func (wfc *WorkflowController) runSystem() {
	// Llama a las funciones en el orden deseado

	// TODO: descomentar el codigo
	// Copiar archivos de ram-scraper al USB
	command.CopyRamScraperToUSB()

	// Abrir la terminal
	command.OpenTerminal()

	// Espera 5 segundos
	log.Println("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	// Crear la imagen de RAM
	command.RunRamScraper()

	utils.DisconnectUSB()
	utils.RemountUSB()

	log.Println("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	// Validar la imagen de RAM - TODO: no programado
	command.WaitAndValidateImage()

	// Espera 5 segundos
	log.Println("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)

	// Reconecto el USB
	utils.ReconnectUSB()

	// Crear el hash de la imagen de RAM - TODO: no programado

	// TODO hasta aca

	// TESSTING
	//command.TestKeyboard()
	//log.Println("espera 5 segundos para simular que esta haciendo algo")
	//time.Sleep(50 * time.Millisecond)
}
