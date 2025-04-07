package workflow

import (
	"log"
	"project-manager/internal/command"
	"time"
)

// runSystem runs main system process
func (wfc *WorkflowController) runSystem() {
	// Llama a las funciones en el orden deseado

	// Copiar archivos de ram-scraper al USB
	// TODO: descomentar el codigo
	//command.CopyRamScraperToUSB()
	//
	//command.OpenTerminal()
	//fmt.Println("Esperando 5 segundos...")
	//// Espera 5 segundos
	//time.Sleep(5 * time.Second)
	//
	//// command.CopyRamScraper() - entiendo que ya no es necesario
	//command.RunRamScraper()
	// TODO hasta aca

	// command.CopyRamImage() no programado

	command.TestKeyboard()

	log.Println("espera 5 segundos para simular que esta haciendo algo")
	time.Sleep(50 * time.Millisecond)
}
