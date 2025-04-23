package files

import (
	"data-seal/internal/logs"
	"os"
)

func WriteToFile(filePath, data string) error {
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	return nil
}
