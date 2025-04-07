package files

import (
	"fmt"
	"os"
)

func WriteToFile(filePath, data string) error {
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return err
	}
	return nil
}
