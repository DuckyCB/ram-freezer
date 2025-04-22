package command

import (
	"log"
	"ram-scraper/pkg/utils"
	"time"
	"ram-scraper/utils/constants"
	"os"
)

func WaitForImageCompletion(waitTime int) int {
	utils.RemountUSB()

	waitTimeSec := time.Duration(waitTime) * time.Second // segundos

	// Cargar la configuracion
	config, err := utils.LoadConfig(constants.ConfigPath)
	if err != nil {
		log.Println("Error cargando la configuracion:", err)
		os.Exit(1)
	}

	// State file path
	stateFilePath := config.StateFile
	// Le hago join con /mnt/usb/
	stateFilePath = "/mnt/usb/" + stateFilePath

	// Cargar el estado
	state, err := utils.LoadState(stateFilePath)
	if err != nil {
		log.Println("Error cargando el estado:", err)
		os.Exit(1)
	}


	// Esperar el tiempo especificado
	for state.Status != "completed" && state.Status != "error"{
		utils.RemountUSB()

		// Cargar el estado
		state, err := utils.LoadState(stateFilePath)
		if err != nil {
			log.Println("Error cargando el estado:", err)
			return 1
		}
	
		if state.Status == "completed"{
			log.Println("La imagen de RAM se ha creado correctamente.")
			return 0
		} else if state.Status == "error" {
			log.Println("Error al crear la imagen de RAM:", *state.ErrorMessage)
			return 1
		} else {
			log.Println("La imagen de RAM no se ha creado. Estado actual:", state.Status)
		}
		log.Printf("Esperando %v\n", waitTimeSec)
		time.Sleep(waitTimeSec)
	}
	return 1
}