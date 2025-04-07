package gpio

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkGPIOAccess() bool {
	_, err := os.Stat("/sys/class/gpio")
	return err == nil
}

func isPinExported(pin int) bool {
	_, err := os.Stat(fmt.Sprintf("/sys/class/gpio/gpio%d", pin))
	return err == nil
}

func getGPIOPin(pin int) (int, error) {
	return pin + 512, nil
}

func initGPIO(pin int, direction string) error {
	if !isPinExported(pin) {
		err := os.WriteFile("/sys/class/gpio/export", []byte(strconv.Itoa(pin)), 0644)
		if err != nil {
			if !strings.Contains(err.Error(), "device or resource busy") {
				return fmt.Errorf("error al exportar pin %d: %v", pin, err)
			}
		}

		time.Sleep(200 * time.Millisecond) // Espera a que se creen los archivos necesarios
	}

	err := os.WriteFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", pin), []byte(direction), 0644)
	if err != nil {
		return fmt.Errorf("error al configurar dirección del pin %d: %v", pin, err)
	}

	return nil
}

func cleanupGPIO(pin int) error {
	err := os.WriteFile("/sys/class/gpio/unexport", []byte(strconv.Itoa(pin)), 0644)
	if err != nil {
		return fmt.Errorf("error al liberar pin %d: %v", pin, err)
	}
	return nil
}

func readGPIO(pin int) (int, error) {
	data, err := os.ReadFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin))
	if err != nil {
		return -1, fmt.Errorf("error al leer pin %d: %v", pin, err)
	}
	value, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return -1, fmt.Errorf("error al convertir valor del pin %d: %v", pin, err)
	}
	return value, nil
}

func writeGPIO(pin int, value int) error {
	err := os.WriteFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin), []byte(strconv.Itoa(value)), 0644)
	if err != nil {
		return fmt.Errorf("error al escribir pin %d: %v", pin, err)
	}
	return nil
}
