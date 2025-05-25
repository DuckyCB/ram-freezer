package system

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"project-manager/internal/logs"
	"ram-freezer/utils/rfutils/pkg/rfutils"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	Serial         string `json:"serial"`
	Version        string `json:"version"`
	GoVersion      string `json:"go_version"`
	BashVersion    string `json:"bash_version"`
	KernelVersion  string `json:"kernel_version"`
	Timezone       string `json:"timezone"`
	USBStorageSize string `json:"usb_storage_size"`
}

// StartRun inicia los directorios para recibir los datos de ejecución
func StartRun() error {
	runPath, err := setNewRunPath()

	systemInfo, err := collectSystemInfo()
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	jsonData, err := json.MarshalIndent(systemInfo, "", "  ")
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	jsonFilePath := filepath.Join(runPath, "sistema.info")
	if err := os.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	logs.Log.Info(fmt.Sprintf("Archivo '%s' creado con éxito.\n", jsonFilePath))

	binaryPath := "/opt/ram-freezer/data-seal"
	logs.Log.Info(fmt.Sprintf("Ejecutando '%s system'...\n", binaryPath))

	cmd := exec.Command(binaryPath, "system")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	logs.Log.Info("Ejecución del binario completada con éxito.")

	return nil
}

func setNewRunPath() (string, error) {
	baseDir := "/opt/ram-freezer/bin"

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		logs.Log.Error(fmt.Sprintf("Error al asegurar el directorio base %s: %s", baseDir, err.Error()))
		return "", err
	}

	nextRunNumber, err := getNextRunNumber(baseDir)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error al obtener el siguiente número de run: %s", err.Error()))
		return "", err
	}

	today := time.Now().Format("2006-01-02")
	newRunDirName := fmt.Sprintf("%s_%d", today, nextRunNumber)
	newRunPath := filepath.Join(baseDir, newRunDirName)

	if err := os.Mkdir(newRunPath, 0755); err != nil {
		logs.Log.Error(fmt.Sprintf("Error al crear el directorio de run '%s': %s", newRunPath, err.Error()))
		return "", err
	}
	logs.Log.Info(fmt.Sprintf("Directorio de run '%s' creado con éxito.", newRunPath))

	outputPath := "/opt/ram-freezer/.out"
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		logs.Log.Error(fmt.Sprintf("Error al crear el directorio para .out: %s", err.Error()))
		return "", err
	}

	err = os.WriteFile(outputPath, []byte(newRunPath+"\n"), 0644)
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error al guardar la ruta de la run en '%s': %s", outputPath, err.Error()))
		return "", err
	}
	logs.Log.Info(fmt.Sprintf("Ruta de la run '%s' guardada en '%s'.", newRunPath, outputPath))

	return newRunPath, nil
}

// getNextRunNumber busca el número de la última carpeta de run creada con el formato YYYY-MM-DD_X
func getNextRunNumber(parentDir string) (int, error) {
	files, err := os.ReadDir(parentDir)
	if err != nil {
		return 1, fmt.Errorf("error al leer el directorio %s: %w", parentDir, err)
	}

	maxNumber := 0
	todayPrefix := time.Now().Format("2006-01-02") + "_"

	for _, file := range files {
		if file.IsDir() {
			fileName := file.Name()
			if strings.HasPrefix(fileName, todayPrefix) {
				parts := strings.Split(fileName, "_")
				if len(parts) == 2 {
					numStr := parts[1]
					num, err := strconv.Atoi(numStr)
					if err == nil {
						if num > maxNumber {
							maxNumber = num
						}
					}
				}
			}
		}
	}
	return maxNumber + 1, nil
}

// collectSystemInfo recopila la información del sistema
func collectSystemInfo() (Info, error) {
	var info Info

	info.Serial = rfutils.GetRaspberryPiSerial()
	info.GoVersion = rfutils.GetGoVersion()
	info.BashVersion = rfutils.GetBashVersion()
	info.Version = rfutils.GetVersion()
	info.KernelVersion = rfutils.GetKernelVersion()
	info.Timezone = time.Now().Location().String()
	info.USBStorageSize = rfutils.GetStorageSize()

	return info, nil
}
