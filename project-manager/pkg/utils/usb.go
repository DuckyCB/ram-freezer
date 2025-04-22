package utils

import (
	"fmt"
	"os/exec"
)

// DisconnectUSB disconnects the USB device by writing to the UDC file
func DisconnectUSB() {
	fmt.Println("Desconectando el USB...")
	// Desconecto el USB
	cmd := exec.Command("bash", "-c", "echo '' | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error reconectando el USB:", string(output))
		return
	}
}

// ConnectUSB connects the USB device by writing to the UDC file
func ConnectUSB() {
	fmt.Println("Conectando el USB...")
	cmd := exec.Command("bash", "-c", "ls /sys/class/udc | sudo tee /sys/kernel/config/usb_gadget/ram-freezer/UDC")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error conectando el USB:", string(output))
		return
	}
}


// MountUSB mounts the USB device to the specified mount point
func MountUSB() {
	fmt.Println("Montando el USB...")
	cmd := exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error montando el USB:", string(output))
		return
	}
	fmt.Println("USB montado correctamente.")
}

// UmountUSB unmounts the USB device from the specified mount point
func UmountUSB() {
	fmt.Println("Desmontando el USB...")
	cmd := exec.Command("sudo", "umount", "/mnt/usb/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error desmontando el USB:", string(output))
		return
	}
	fmt.Println("USB desmontado correctamente.")
}

// RemountUSB remounts the USB device to ensure it is properly connected
func RemountUSB() {
	fmt.Println("Remontando el USB...")
	UmountUSB()
	MountUSB()
	fmt.Println("USB remontado correctamente.")
}

// ReconnectUSB reconnects the USB device by first disconnecting and then connecting it
func ReconnectUSB() {
	fmt.Println("Reconectando el USB...")
	DisconnectUSB()
	UmountUSB()
	MountUSB()
	ConnectUSB()
	fmt.Println("USB reconectado correctamente.")
}