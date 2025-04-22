package command

import (
	"fmt"
	"os"
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
		fmt.Printf("ERROR: No se pudo cargar la configuracion: %v\n", err)
		os.Exit(1)
	}
	// State file path
	stateFilePath := config.StateFile
	// Le hago join con /mnt/usb/
	stateFilePath = "/mnt/usb/" + stateFilePath
	
	state, err := utils.LoadState(stateFilePath)
	fmt.log.Println("Cargando el estado desde:", stateFilePath)
	if err != nil {
		fmt.Printf("ERROR: No se pudo cargar el estado: %v\n", err)
		os.Exit(1)
	}

	filePath := config.OutputFolder + "ps1/" + config.OutputFile
	fmt.Printf("Validando el archivo %s...\n", filePath)

	// Verificar si el archivo existe
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		msg := fmt.Sprintf("ERROR: El archivo %s no fue generado.\n", filePath)
		exit_code := 1
		utils.WriteStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}

	// Obtener el tamaño del archivo en bytes
	fileSize := fileInfo.Size()
	// Convertir a GB
	fileSizeGB := bytesToGB(fileSize)

	// Obtener la cantidad total de RAM del sistema
	totalRAMGB := state.TotalRAM
	
	fmt.Printf("Tamaño total de RAM: %.2f GB\n", totalRAMGB)

	// Definir un margen de error del 5%
	lowerLimit := totalRAMGB * 0.95 // 95% de la RAM total
	upperLimit := totalRAMGB * 1.2 // 120% de la RAM total

	fmt.Printf("Tamaño del archivo: %.2f GB\n", fileSizeGB)
	fmt.Printf("Límite inferior: %.2f GB\n", lowerLimit)
	fmt.Printf("Límite superior: %.2f GB\n", upperLimit)


	// Verificar si el archivo es demasiado pequeño
	if fileSizeGB < 1 {
		msg := fmt.Sprintf("ERROR: El archivo %s es demasiado pequeño para ser válido. Tamaño: %.2f GB\n", filePath, fileSizeGB)
		exit_code := 1
		utils.WriteStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}

	// Verificar si el archivo tiene un tamaño adecuado en comparación con la RAM
	if fileSizeGB < lowerLimit {
		msg := fmt.Sprintf("WARNING: El archivo %s es más pequeño de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
		exit_code := 0
		utils.WriteStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	} else if fileSizeGB > upperLimit {
		msg := fmt.Sprintf("WARNING: El archivo %s es más grande de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
		exit_code := 0
		utils.WriteStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}

	// Intentar abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("ERROR: No se pudo abrir el archivo %s. Puede estar corrupto.", filePath)
		exit_code := 1
		utils.WriteStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}
	defer file.Close()

	// Si todo está bien
	msg := fmt.Sprintf("SUCCESS: El archivo %s es válido. Tamaño: %.2f GB, memoria RAM fisica: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
	exit_code := 0
	utils.WriteStateVal(config.StateFile, state, msg, exit_code)
	fmt.Printf(msg)
	os.Exit(exit_code)
}
