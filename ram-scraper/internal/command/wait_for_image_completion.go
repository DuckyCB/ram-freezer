package command

import (
	"log"
	"ram-scraper/pkg/utils"
	"time"
	"ram-scraper/utils/constants"
	"os"
	"os/exec"
)

func remountUSB() {
	cmd := exec.Command("bash", "-c", "mount | grep '/dev/sda1'")
	output, err := cmd.CombinedOutput()
	if err != nil || len(output) == 0 {
		log.Println("El USB no está montado. Intentando montarlo...")
		cmd := exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
		output, err = cmd.CombinedOutput()
		if err != nil {
			log.Println("Error montando el USB:", string(output))
		}
		log.Println("USB montado correctamente.")
	}
	//log.Println("El USB ya está montado. Desmontando y volviendo a montar...")
	cmd = exec.Command("bash", "-c", "sudo umount /dev/sda1 && sudo mount /dev/sda1 /mnt/usb/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Println("Error desmontando y volviendo a montar el USB:", string(output))
		return 
	}
	//log.Println("USB desmontado y montado correctamente.")
}

func WaitForImageCompletion(waitTime int) int {
	remountUSB()

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
		remountUSB()

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