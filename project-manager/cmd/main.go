package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
			CopyRamScraper()
			RunRamScraper()
			CopyRamImage()
		}
	}
}
