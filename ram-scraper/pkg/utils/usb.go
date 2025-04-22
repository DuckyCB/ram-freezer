package utils

import (
	"log"
	"os/exec"
)

func RemountUSB() {
	cmd := exec.Command("bash", "-c", "mount | grep '/dev/sda1'")
	output, err := cmd.CombinedOutput()
	if err != nil || len(output) == 0 {
		log.Println("El USB no está montado. Intentando montarlo...")
		cmd := exec.Command("sudo", "mount", "/dev/sda1", "/mnt/usb/")
		output, err = cmd.CombinedOutput()
		if err != nil {
			log.Println("Error montando el USB:", string(output))
		}
		log.Println("USB montado correctamente.")
	}
	//log.Println("El USB ya está montado. Desmontando y volviendo a montar...")
	cmd = exec.Command("bash", "-c", "sudo umount /dev/sda1 && sudo mount /dev/sda1 /mnt/usb/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Println("Error desmontando y volviendo a montar el USB:", string(output))
		return
	}
	//log.Println("USB desmontado y montado correctamente.")
}
