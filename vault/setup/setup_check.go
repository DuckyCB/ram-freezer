package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// checkUsbGadgetService verifica si el servicio usb-gadget está activo en systemd
func checkUsbGadgetService() bool {
	cmd := exec.Command("systemctl", "is-active", "usb-gadget")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	status := strings.TrimSpace(string(output))
	return status == "active" || status == "activating"
}

// checkModuleLoaded verifica si un módulo del kernel está cargado
func checkModuleLoaded(moduleName string) bool {
	out, err := exec.Command("lsmod").Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(out), moduleName)
}

// checkMountpoint verifica si un directorio está montado
func checkMountpoint(mountPoint string) bool {
	data, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return false
	}
	return strings.Contains(string(data), mountPoint)
}

// checkConfigFSContent verifica si hay contenido en ConfigFS (que no esté vacío)
func checkConfigFSContent() bool {
	files, err := os.ReadDir("/sys/kernel/config")
	return err == nil && len(files) > 0
}

// checkStorageDevice verifica si existe un medio de almacenamiento
func checkStorageDevice() bool {
	_, err := os.Stat("/dev/sda1")
	return err == nil
}

// isExFAT verifies that the usb drive uses exfat as file system
func isExFAT() bool {
	cmd := exec.Command("blkid", "/dev/sda1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.Contains(string(output), "TYPE=\"exfat\"")
}

func getAvailableSpaceGB() float64 {
	cmd := exec.Command("df", "-P", "/dev/sda1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return 0
	}
	fields := strings.Fields(lines[1])
	if len(fields) < 4 {
		return 0
	}
	availableSpaceKBStr := fields[3]
	availableSpaceKB, err := strconv.ParseUint(availableSpaceKBStr, 10, 64)
	if err != nil {
		return 0
	}
	return float64(availableSpaceKB) / (1024 * 1024)
}

func checkDiskLabel() bool {
	cmd := exec.Command("blkid", "/dev/sda1")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	output := string(outputBytes)
	return strings.Contains(output, fmt.Sprintf("LABEL=\"%s\"", "USB_VAULT"))
}

func main() {
	fmt.Println("📦🔒 Verificando Vault 📦🔒")

	availableSpaceGB := getAvailableSpaceGB()

	checks := []struct {
		name   string
		result bool
	}{
		{" - usb-gadget.service está activo", checkUsbGadgetService()},
		{" - Módulo dwc2 cargado", checkModuleLoaded("dwc2")},
		{" - Módulo libcomposite cargado", checkModuleLoaded("libcomposite")},
		{" - ConfigFS montado", checkMountpoint("/sys/kernel/config")},
		{" - ConfigFS tiene contenido", checkConfigFSContent()},
		{" - Almacenamiento existente", checkStorageDevice()},
		{" - El tipo es exFAT", isExFAT()},
		{fmt.Sprintf(" - Quedan %.2f GB disponibles", availableSpaceGB), availableSpaceGB >= 4},
		{" - La etiqueta es USB_VAULT", checkDiskLabel()},
	}

	fmt.Println("📋 Resultados de la instalación de vault:")
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
		fmt.Println("🎉 Todo está correctamente configurado en vault.")
	} else {
		fmt.Println("⚠️ Hay problemas en la configuración de vault. Revisa los errores.")
	}
}
