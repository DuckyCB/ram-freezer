package command

import (
	"fmt"
	"os"
	"ram-scraper/internal/logs"
	"ram-scraper/pkg/utils"
	"ram-scraper/utils/constants"
)

// bytesToGB convierte bytes a GB con dos decimales
func bytesToGB(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func ValidateOutput() {
	config, err := utils.LoadConfig(constants.ConfigPath)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("No se pudo cargar la configuracion: %v\n", err))
		os.Exit(1)
	}
	usbFolder := "/mnt/usb/"
	stateFilePath := config.StateFile
	stateFilePath = usbFolder + stateFilePath

	state, err := utils.LoadState(stateFilePath)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("No se pudo cargar el estado: %v", err))
		os.Exit(1)
	}

	filePath := usbFolder + config.OutputFolder + config.OutputFile
	logs.Log.Info(fmt.Sprintf("Validando el archivo %s...", filePath))

	// Verificar si el archivo existe
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		msg := fmt.Sprintf("El archivo %s no fue generado.", filePath)
		exit_code := 1
		utils.WriteStateVal(stateFilePath, state, msg, exit_code)
		logs.Log.Error(msg)
		os.Exit(exit_code)
	}

	// Obtener el tamaño del archivo en bytes
	fileSize := fileInfo.Size()
	// Convertir a GB
	fileSizeGB := bytesToGB(fileSize)

	// Obtener la cantidad total de RAM del sistema
	totalRAMGB := state.TotalRAM

	logs.Log.Info(fmt.Sprintf("Tamaño total de RAM: %.2f GB", totalRAMGB))

	// Definir un margen de error del 5%
	lowerLimit := totalRAMGB * 0.95 // 95% de la RAM total
	upperLimit := totalRAMGB * 1.2  // 120% de la RAM total

	logs.Log.Info(fmt.Sprintf("Tamaño del archivo: %.2f GB", fileSizeGB))
	logs.Log.Info(fmt.Sprintf("Límite inferior: %.2f GB", lowerLimit))
	logs.Log.Info(fmt.Sprintf("Límite superior: %.2f GB", upperLimit))

	// Verificar si el archivo es demasiado pequeño
	if fileSizeGB < 1 {
		msg := fmt.Sprintf("El archivo %s es demasiado pequeño para ser válido. Tamaño: %.2f GB", filePath, fileSizeGB)
		exit_code := 1
		utils.WriteStateVal(stateFilePath, state, msg, exit_code)
		logs.Log.Error(msg)
		os.Exit(exit_code)
	}

	// Verificar si el archivo tiene un tamaño adecuado en comparación con la RAM
	if fileSizeGB < lowerLimit {
		msg := fmt.Sprintf("El archivo %s es más pequeño de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB", filePath, fileSizeGB, totalRAMGB)
		exit_code := 0
		utils.WriteStateVal(stateFilePath, state, msg, exit_code)
		logs.Log.Warn(msg)
		os.Exit(exit_code)
	} else if fileSizeGB > upperLimit {
		msg := fmt.Sprintf("El archivo %s es más grande de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB", filePath, fileSizeGB, totalRAMGB)
		exit_code := 0
		utils.WriteStateVal(stateFilePath, state, msg, exit_code)
		logs.Log.Warn(msg)
		os.Exit(exit_code)
	}

	// Intentar abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("No se pudo abrir el archivo %s. Puede estar corrupto.", filePath)
		exit_code := 1
		utils.WriteStateVal(stateFilePath, state, msg, exit_code)
		logs.Log.Error(msg)
		os.Exit(exit_code)
	}
	defer file.Close()

	msg := fmt.Sprintf("El archivo %s es válido. Tamaño: %.2f GB, memoria RAM fisica: %.2f GB", filePath, fileSizeGB, totalRAMGB)
	exit_code := 0
	utils.WriteStateVal(stateFilePath, state, msg, exit_code)
	logs.Log.Info(msg)

	os.Exit(exit_code)
}
