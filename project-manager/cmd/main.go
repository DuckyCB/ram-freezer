package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"project-manager/internal/workflow"
	"project-manager/pkg/utils"
	"project-manager/utils/constants"
	"syscall"
)

func main() {
	log.Println("Starting project manager")

	if !utils.IsAdmin() {
		log.Println("Es necesario ejecutar el programa como administrador")
		return
	}

	controller, err := workflow.NewWorkflowController(constants.LedPin, constants.ButtonPin)
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
