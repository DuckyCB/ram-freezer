package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"project-manager/cmd/utils"
	"syscall"
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
	log.Println("espera 5 segundos para simular que esta haciendo algo")
	time.Sleep(50 * time.Millisecond)
}

func main() {
	log.Println("Starting project manager")

	if !utils.IsAdmin() {
		log.Println("Es necesario ejecutar el programa como administrador")
		return
	}

	controller, err := NewWorkflowController(ledPin, buttonPin)
	if err != nil {
		fmt.Printf("Error al inicializar el sistema: %v\n", err)
		return
	}
	defer controller.Stop()

	controller.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
