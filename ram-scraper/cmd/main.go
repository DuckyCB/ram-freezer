package main

import (
	"ram-scraper/internal/command"
	"ram-scraper/internal/logs"
	"ram-scraper/utils/constants"
)

func main() {
	logs.SetupLogger()

	logs.Log.Info("Esperando la creacion de la imagen RAM...")
	command.WaitForImageCompletion(constants.WaitTime)
	logs.Log.Info("La imagen de RAM se ha creado.")

	logs.Log.Info("Validando la imagen de RAM...")
	command.ValidateOutput()

	logs.Log.Info("La imagen de RAM ha sido validada.")
}
