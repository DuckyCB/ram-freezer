package utils

import (
	"fmt"
	"os/exec"
	"ram-scraper/internal/logs"
)

func RemountUSB() {
	cmd := exec.Command("bash", "-c", "mount | grep '/dev/sda1'")
	output, err := cmd.CombinedOutput()
	if err != nil || len(output) == 0 {
		logs.Log.Info("El USB no está montado. Intentando montarlo...")
		cmd := exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
		output, err = cmd.CombinedOutput()
		if err != nil {
			logs.Log.Error(fmt.Sprintf("Error montando el USB: %v", string(output)))
		}
		logs.Log.Info("USB montado correctamente.")
	}
	//log.Println("El USB ya está montado. Desmontando y volviendo a montar...")
	cmd = exec.Command("bash", "-c", "sudo umount /dev/sda1 && sudo mount /dev/sda1 /mnt/usb/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		logs.Log.Info(fmt.Sprintf("Error desmontando y volviendo a montar el USB: %s", string(output)))
		return
	}
	//log.Println("USB desmontado y montado correctamente.")
}
