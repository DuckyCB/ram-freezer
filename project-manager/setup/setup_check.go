package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// checkProjectManagerService verifica si el servicio project-manager está activo en systemd
func checkProjectManagerService() bool {
	cmd := exec.Command("systemctl", "status", "project-manager")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return false
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return false
	}

	if err := cmd.Start(); err != nil {
		return false
	}

	scanner := bufio.NewScanner(stdout)
	state := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Active:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				state = strings.TrimSpace(parts[1])
				break
			}
		}
	}

	errStderr := ""
	errScanner := bufio.NewScanner(stderr)
	for errScanner.Scan() {
		errStderr += errScanner.Text() + "\n"
	}
	if err := errScanner.Err(); err != nil {
		return strings.Contains(state, "active") || strings.Contains(state, "activating")
	}

	if err := cmd.Wait(); err != nil {
		return strings.Contains(state, "active") || strings.Contains(state, "activating")
	}

	return strings.Contains(state, "active") || strings.Contains(state, "activating")
}

func isPinExported(pin int) bool {
	_, err := os.Stat(fmt.Sprintf("/sys/class/gpio/gpio%d", pin))
	return err == nil
}

func isGPIODirection(pin int, expectedDirection string) bool {
	directionFile := fmt.Sprintf("/sys/class/gpio/gpio%d/direction", pin)

	contentBytes, err := os.ReadFile(directionFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("GPIO pin %d is not exported", pin)
			return false
		}
		fmt.Printf("error reading direction for GPIO pin %d: %v", pin, err)
		return false
	}

	direction := strings.TrimSpace(string(contentBytes))
	return direction == expectedDirection
}

func main() {
	fmt.Println("🧙 Project Manager 🧙")

	checks := []struct {
		name   string
		result bool
	}{
		{" - project-manager.service está activo", checkProjectManagerService()},
		{" - pin 17 exportado", isPinExported(529)},          // (512 + 17 = 529)
		{" - pin 29 exportado", isPinExported(539)},          // (512 + 27 = 539)
		{" - pin 17 es input", isGPIODirection(529, "in")},   // (512 + 17 = 529)
		{" - pin 29 es output", isGPIODirection(539, "out")}, // (512 + 27 = 539)
	}

	fmt.Println("📋 Resultados de la instalación de project manager:")
	for _, check := range checks {
		status := "✅ OK"
		if !check.result {
			status = "❌ FAIL"
		}
		fmt.Printf("%s: %s\n", check.name, status)
	}

	allOk := true
	for _, check := range checks {
		if !check.result {
			allOk = false
			break
		}
	}

	if allOk {
		fmt.Println("🎉 Todo está correctamente configurado en project manager.")
	} else {
		fmt.Println("⚠️ Hay problemas en la configuración de project manager. Revisa los errores.")
	}
}
