package workflow

import (
	"log"
	"project-manager/internal/command"
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
	fmt.Println("Esperando 5 segundos...")
	time.Sleep(5 * time.Second)
	

	// Crear la imagen de RAM
	command.RunRamScraper()
	
	// Validar la imagen de RAM - TODO: no programado

	// Crear el hash de la imagen de RAM - TODO: no programado

	// TODO hasta aca

	command.TestKeyboard()

	//log.Println("espera 5 segundos para simular que esta haciendo algo")
	//time.Sleep(50 * time.Millisecond)
}
