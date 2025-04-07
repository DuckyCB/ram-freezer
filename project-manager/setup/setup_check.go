package main

import (
	"bufio"
	"fmt"
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

func main() {
	fmt.Println("🧙 Project Manager 🧙")

	checks := []struct {
		name   string
		result bool
	}{
		{" - project-manager.service está activo", checkProjectManagerService()},
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
