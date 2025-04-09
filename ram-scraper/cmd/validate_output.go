package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de configuracion
type Config struct {
	ExeName      string `json:"exe_name"`
	OutputFile   string `json:"output_file"`
	OutputFolder string `json:"output_folder"`
	LogFolder    string `json:"log_folder"`
	LogFile      string `json:"log_file"`
	StateFile    string `json:"state_file"`
}

type State struct {
	Status       string   `json:"status"`
	StartTime    string   `json:"start_time"`
	EndTime      string   `json:"end_time"`
	Duration     float64  `json:"duration"`
	ErrorMessage *string  `json:"error_message"` // puntero para soportar null
	TotalRAM     float64   `json:"total_ram"`
	ValidationMessage string `json:"validation_message"` // puntero para soportar null
	ValidationExitCode int `json:"validation_exit_code"`
}

func LoadState(path string) (*State, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error leyendo state.json: %w", err)
	}

	var state State
	if err := json.Unmarshal(bytes, &state); err != nil {
		return nil, fmt.Errorf("error parseando state.json: %w", err)
	}

	return &state, nil
}

func LoadConfig(path string) (*Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error leyendo state.json: %w", err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("error parseando config.json: %w", err)
	}

	return &config, nil
}


// bytesToGB convierte bytes a GB con dos decimales
func bytesToGB(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

// escribir en state el estado de la validacion
func writeStateVal(path string, state *State, val_msg string, val_exit_code int) error {

	state.ValidationMessage = val_msg
	state.Status = "VALIDATION"
	state.ValidationExitCode = val_exit_code

	bytes, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar el estado: %w", err)
	}

	if err := os.WriteFile(path, bytes, 0644); err != nil {
		return fmt.Errorf("error al escribir el estado: %w", err)
	}

	return nil
}

func main() {
	config, err := LoadConfig("config/settings.json")
	if err != nil {
		fmt.Printf("ERROR: No se pudo cargar la configuracion: %v\n", err)
		os.Exit(1)
	}
	
	state, err := LoadState(config.StateFile)
	if err != nil {
		fmt.Printf("ERROR: No se pudo cargar el estado: %v\n", err)
		os.Exit(1)
	}

	currentDir, err := os.Getwd()

	filePath := currentDir + "/" + config.OutputFolder + "ps1/" + config.OutputFile
	fmt.Printf("Validando el archivo %s...\n", filePath)

	// Verificar si el archivo existe
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		msg := fmt.Sprintf("ERROR: El archivo %s no fue generado.\n", filePath)
		exit_code := 1
		writeStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
		os.Exit(1)
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
		writeStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}

	// Verificar si el archivo tiene un tamaño adecuado en comparación con la RAM
	if fileSizeGB < lowerLimit {
		msg := fmt.Sprintf("WARNING: El archivo %s es más pequeño de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
		exit_code := 0
		writeStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	} else if fileSizeGB > upperLimit {
		msg := fmt.Sprintf("WARNING: El archivo %s es más grande de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
		exit_code := 0
		writeStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}

	// Intentar abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("ERROR: No se pudo abrir el archivo %s. Puede estar corrupto.", filePath)
		exit_code := 1
		writeStateVal(config.StateFile, state, msg, exit_code)
		fmt.Printf(msg)
		os.Exit(exit_code)
	}
	defer file.Close()

	// Si todo está bien
	msg := fmt.Sprintf("SUCCESS: El archivo %s es válido. Tamaño: %.2f GB, memoria RAM fisica: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
	exit_code := 0
	writeStateVal(config.StateFile, state, msg, exit_code)
	fmt.Printf(msg)
	os.Exit(exit_code)
}
