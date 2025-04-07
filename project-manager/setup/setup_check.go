package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// checkUsbGadgetService verifica si el servicio usb-gadget está activo en systemd
func checkUsbGadgetService() bool {
	cmd := exec.Command("systemctl", "is-active", "project-manager")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	status := strings.TrimSpace(string(output))
	return status == "active" || status == "activating"
}

func main() {
	fmt.Println("🧙 Project Manager 🧙")

	checks := []struct {
		name   string
		result bool
	}{
		{" - project-manager.service está activo", checkUsbGadgetService()},
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
		fmt.Println("\n🎉 Todo está correctamente configurado en project manager.")
	} else {
		fmt.Println("\n⚠️ Hay problemas en la configuración de project manager. Revisa los errores.")
	}
}
