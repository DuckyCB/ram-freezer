package main

import (
	"fmt"
	"os"
	"syscall"
)

// levantamos el archivo de configuracion
// de json en config/settings.json
config := Config{}
err := config.Load("config/settings.json")
if err != nil {
	fmt.Println("ERROR: No se pudo cargar el archivo de configuración.")
	os.Exit(1)
}

// getTotalRAM obtiene la cantidad total de RAM del sistema en bytes
func getTotalRAM() (int) {
	// Lo sacamos de la configuracion
	totalRAM := config.TotalRAM
}

// bytesToGB convierte bytes a GB con dos decimales
func bytesToGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func main() {
	// Verificar si se pasó un argumento (ruta del archivo)
	if len(os.Args) < 2 {
		fmt.Println("ERROR: Debes proporcionar la ruta del archivo como argumento.")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Verificar si el archivo existe
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("ERROR: El archivo %s no fue generado.\n", filePath)
		os.Exit(1)
	}

	// Obtener el tamaño del archivo en bytes
	fileSize := fileInfo.Size()

	// Obtener la cantidad total de RAM del sistema
	totalRAM, err := getTotalRAM()
	if err != nil {
		fmt.Println("ERROR: No se pudo obtener la cantidad de RAM del sistema.")
		os.Exit(4)
	}

	// Convertir a GB
	fileSizeGB := bytesToGB(uint64(fileSize))
	totalRAMGB := bytesToGB(totalRAM)

	// Definir un margen de error del 5%
	lowerLimit := uint64(float64(totalRAM) * 0.95) // 95% de la RAM total
	upperLimit := uint64(float64(totalRAM) * 1.05) // 105% de la RAM total

	// Verificar si el archivo es demasiado pequeño
	if fileSize < 1024 {
		fmt.Printf("WARNING: El archivo %s es demasiado pequeño para ser válido. Tamaño: %.2f GB\n", filePath, fileSizeGB)
		os.Exit(2)
	}

	// Verificar si el archivo tiene un tamaño adecuado en comparación con la RAM
	if uint64(fileSize) < lowerLimit {
		fmt.Printf("WARNING: El archivo %s es más pequeño de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
		os.Exit(5)
	} else if uint64(fileSize) > upperLimit {
		fmt.Printf("WARNING: El archivo %s es más grande de lo esperado. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
		os.Exit(6)
	}

	// Intentar abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("WARNING: No se pudo abrir el archivo %s. Puede estar corrupto.\n", filePath)
		os.Exit(3)
	}
	defer file.Close()

	// Si todo está bien
	fmt.Printf("SUCCESS: El archivo %s es válido. Tamaño: %.2f GB, RAM esperada: %.2f GB\n", filePath, fileSizeGB, totalRAMGB)
	os.Exit(0)
}
