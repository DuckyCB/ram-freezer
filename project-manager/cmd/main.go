package main

import (
	"fmt"
	"os"
	"os/signal"
	"project-manager/internal/logs"
	"project-manager/internal/system"
	"project-manager/internal/workflow"
	"project-manager/pkg/utils"
	"project-manager/utils/constants"
	"syscall"
)

func main() {
	logs.SetupLogger()

	system.StartRun()

	logs.Log.Info("Starting project manager")

	if !utils.IsAdmin() {
		logs.Log.Fatal("Es necesario ejecutar el programa como administrador")
		return
	}

	controller, err := workflow.NewController(constants.LedPin, constants.ButtonPin)
	if err != nil {
		logs.Log.Fatal(fmt.Sprintf("Error al inicializar el sistema: %v", err))
		return
	}
	defer controller.Stop()

	controller.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
