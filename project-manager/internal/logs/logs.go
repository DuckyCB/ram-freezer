package logs

import (
	"fmt"
	"ram-freezer/audit-trail/pkg/logger"
)

var Log *logger.SimpleLogger

func SetupLogger() {
	var err error

	Log, err = logger.NewRFLogger()
	if err != nil {
		fmt.Println("Error al inicializar el logger:", err)
	}
}
