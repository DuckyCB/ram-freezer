package command

import (
	"fmt"
	"os"
	"ram-scraper/internal/logs"
	"ram-scraper/pkg/utils"
	"ram-scraper/utils/constants"
	"time"
)

func WaitForImageCompletion(waitTime int) int {
	utils.RemountUSB()

	waitTimeSec := time.Duration(waitTime) * time.Second // segundos

	// Cargar la configuracion
	config, err := utils.LoadConfig(constants.ConfigPath)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error cargando la configuracion: %v", err))
		os.Exit(1)
	}

	// State file path
	stateFilePath := config.StateFile
	// Le hago join con /mnt/usb/
	stateFilePath = "/mnt/usb/" + stateFilePath

	// Cargar el estado
	state, err := utils.LoadState(stateFilePath)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error cargando el estado: %v", err))
		os.Exit(1)
	}

	// Esperar el tiempo especificado
	for state.Status != "completed" && state.Status != "error" {
		utils.RemountUSB()

		// Cargar el estado
		state, err := utils.LoadState(stateFilePath)
		if err != nil {
			logs.Log.Error(fmt.Sprintf("Error cargando el estado: %v", err))
			return 1
		}

		if state.Status == "completed" {
			logs.Log.Info("La imagen de RAM se ha creado correctamente.")
			return 0
		} else if state.Status == "error" {
			logs.Log.Error(fmt.Sprintf("Error al crear la imagen de RAM: %v", *state.ErrorMessage))
			return 1
		} else {
			logs.Log.Info(fmt.Sprintf("La imagen de RAM no se ha creado. Estado actual: %s", state.Status))
		}
		logs.Log.Info(fmt.Sprintf("Esperando %v", waitTimeSec))
		time.Sleep(waitTimeSec)
	}
	return 1
}
