package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

// checkDeviceExists verifica si un archivo/dispositivo existe
func checkDeviceExists(devicePath string) bool {
	_, err := os.Stat(devicePath)
	return err == nil
}

// checkPermissions verifica si el dispositivo tiene permisos de escritura
func checkPermissions(devicePath string) bool {
	info, err := os.Stat(devicePath)
	if err != nil {
		return false
	}
	mode := info.Mode()
	return mode&0222 != 0 // Permisos de escritura para usuario, grupo u otros
}

// checkConfigFSContent verifica si hay contenido en ConfigFS (que no esté vacío)
func checkConfigFSContent() bool {
	files, err := os.ReadDir("/sys/kernel/config")
	return err == nil && len(files) > 0
}

// checkUdevSettle verifica si `udevadm settle` ha terminado correctamente
func checkUdevSettle() bool {
	err := exec.Command("udevadm", "settle").Run()
	return err == nil
}

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

func main() {
	fmt.Println("👻⌨ Verificando Ghost Keyboard 👻⌨")

	checks := []struct {
		name   string
		result bool
	}{
		{" - usb-gadget.service está activo", checkUsbGadgetService()},
		{" - Módulo dwc2 cargado", checkModuleLoaded("dwc2")},
		{" - Módulo libcomposite cargado", checkModuleLoaded("libcomposite")},
		{" - ConfigFS montado", checkMountpoint("/sys/kernel/config")},
		{" - ConfigFS tiene contenido", checkConfigFSContent()},
		{" - /dev/hidg0 existe", checkDeviceExists("/dev/hidg0")},
		{" - /dev/hidg0 tiene permisos de escritura", checkPermissions("/dev/hidg0")},
		{" - udev ha finalizado (udevadm settle)", checkUdevSettle()},
	}

	fmt.Println("📋 Resultados de la instalación de ghost keyboard:")
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
		fmt.Println("🎉 Todo está correctamente configurado en ghost keyboard.")
	} else {
		fmt.Println("⚠️ Hay problemas en la configuración de ghost keyboard. Revisa los errores.")
	}
}
