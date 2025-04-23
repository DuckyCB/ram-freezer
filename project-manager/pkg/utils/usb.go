package utils

import (
	"fmt"
	"os/exec"
	"project-manager/internal/logs"
)

// DisconnectUSB disconnects the USB device by writing to the UDC file
func DisconnectUSB() {
	logs.Log.Info("Desconectando el USB...")
	// Desconecto el USB
	cmd := exec.Command("bash", "-c", "echo '' | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("%s. %s", output, err.Error()))
		return
	}
}

// ConnectUSB connects the USB device by writing to the UDC file
func ConnectUSB() {
	logs.Log.Info("Conectando el USB...")
	cmd := exec.Command("bash", "-c", "ls /sys/class/udc | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("%s. %s", output, err.Error()))
		return
	}
}

// MountUSB mounts the USB device to the specified mount point
func MountUSB() {
	logs.Log.Info("Montando el USB...")
	cmd := exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("%s. %s", output, err.Error()))
		return
	}
	logs.Log.Info("USB montado correctamente.")
}

// UmountUSB unmounts the USB device from the specified mount point
func UmountUSB() {
	logs.Log.Info("Desmontando el USB...")
	cmd := exec.Command("sudo", "umount", "/mnt/usb/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Log.Error(fmt.Sprintf("%s. %s", output, err.Error()))
		return
	}
	logs.Log.Info("USB desmontado correctamente.")
}

// RemountUSB remounts the USB device to ensure it is properly connected
func RemountUSB() {
	logs.Log.Info("Remontando el USB...")
	UmountUSB()
	MountUSB()
	logs.Log.Info("USB remontado correctamente.")
}

// ReconnectUSB reconnects the USB device by first disconnecting and then connecting it
func ReconnectUSB() {
	logs.Log.Info("Reconectando el USB...")
	DisconnectUSB()
	UmountUSB()
	MountUSB()
	ConnectUSB()
	logs.Log.Info("USB reconectado correctamente.")
}
