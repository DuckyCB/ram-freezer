package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config Estructura de configuracion
type Config struct {
	ExeName      string `json:"exe_name"`
	OutputFile   string `json:"output_file"`
	OutputFolder string `json:"output_folder"`
	LogFolder    string `json:"log_folder"`
	LogFile      string `json:"log_file"`
	StateFile    string `json:"state_file"`
}

type State struct {
	Status             string  `json:"status"`
	StartTime          string  `json:"start_time"`
	EndTime            string  `json:"end_time"`
	Duration           float64 `json:"duration"`
	ErrorMessage       *string `json:"error_message"` // puntero para soportar null
	TotalRAM           float64 `json:"total_ram"`
	ValidationMessage  string  `json:"validation_message"` // puntero para soportar null
	ValidationExitCode int     `json:"validation_exit_code"`
}

func LoadState(path string) (*State, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error leyendo %s: %w", path, err)
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
		return nil, fmt.Errorf("error leyendo %s: %w", path, err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("error parseando config.json: %w", err)
	}

	return &config, nil
}

// WriteStateVal escribir en state el estado de la validacion
func WriteStateVal(path string, state *State, valMsg string, valExitCode int) error {
	fmt.Printf("Escribiendo en el estado: %s\n", path)
	state.ValidationMessage = valMsg
	state.Status = "validation"
	state.ValidationExitCode = valExitCode

	bytes, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar el estado: %w", err)
	}

	// Abrimos el archivo manualmente para poder hacer f.Sync()
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer f.Close()

	if _, err := f.Write(bytes); err != nil {
		return fmt.Errorf("error al escribir en el archivo: %w", err)
	}

	// Forzar que los cambios se escriban en el dispositivo físico
	if err := f.Sync(); err != nil {
		return fmt.Errorf("error al sincronizar el archivo: %w", err)
	}

	return nil
}
