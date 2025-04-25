package command

import (
	"fmt"
	"os/exec"
	"project-manager/internal/logs"
)

func HashFiles(directorio string) {
	logs.Log.Info("Creando el hash de los archivos en el directorio: " + directorio)

	cmd := exec.Command("/opt/ram-freezer/bin/data-seal", "-dir", directorio)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("Error ejecutando el binario: %v", err))
		return
	}

	logs.Log.Info(string(output))
}
