package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	log.Println("Starting project manager")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Escriba 's' para comenzar: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al leer la entrada:", err)
			break
		}

		if input == "s" {
			OpenTerminal()
			fmt.Println("Esperando 5 segundos...")
			// Espera 5 segundos
			time.Sleep(5 * time.Second)
			CopyRamScraper()
			RunRamScraper()
			CopyRamImage()
		}
	}
}
