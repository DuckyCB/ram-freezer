package main

import (
	"ram-scraper/internal/command"
	"ram-scraper/utils/constants"
	"fmt"
)

func main() {
	fmt.Println("Esperando la creacion de la imagen RAM...")
	command.WaitForImageCompletion(constants.WaitTime)
	fmt.Println("La imagen de RAM se ha creado.")

	fmt.Println("Validando la imagen de RAM...")
	command.ValidateOutput()
}