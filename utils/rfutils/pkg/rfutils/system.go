package rfutils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

func GetOutPath() string {
	content, err := os.ReadFile("/opt/ram-freezer/.out")
	if err != nil {
		fmt.Printf("no se pudo leer la ubicación de salida del sistema")
		return "/opt/ram-freezer/bin"
	}
	return string(content)
}

func GetVersion() string {
	content, err := os.ReadFile("/opt/ram-freezer/.version")
	if err != nil {
		fmt.Printf("No se pudo leer la versión del sistema")
		return "dev"
	}
	return strings.TrimSpace(string(content))
}

func GetKernelVersion() string {
	uname := &syscall.Utsname{}
	if err := syscall.Uname(uname); err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	releaseBytes := make([]byte, len(uname.Release))
	for i, v := range uname.Release {
		releaseBytes[i] = byte(v)
	}
	return string(bytes.TrimRight(releaseBytes, "\x00"))
}

func GetStorageSize() string {
	usbStorageSize, err := getUSBStorageSize("/dev/sda1")
	if err != nil {
		fmt.Printf(err.Error())
		return "N/A"
	}
	return usbStorageSize

}

// getUSBStorageSize intenta obtener el tamaño de un dispositivo de almacenamiento
func getUSBStorageSize(devicePath string) (string, error) {
	cmd := exec.Command("df", "-h", devicePath)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		fmt.Printf(err.Error())
		return "", err
	}

	fields := regexp.MustCompile(`\s+`).Split(lines[1], -1)
	if len(fields) > 1 {
		return fields[1], nil
	}

	err = fmt.Errorf("no se pudo parsear el tamaño de %s", devicePath)
	fmt.Printf(err.Error())
	return "", err
}

func GetRaspberryPiSerial() string {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	fmt.Printf("No se encontró el número de serie")
	return ""
}

func GetGoVersion() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	return strings.TrimSpace(string(output))
}

func GetBashVersion() string {
	cmd := exec.Command("bash", "-c", "echo $BASH_VERSION")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0])
	}
	fmt.Printf("no se pudo obtener la versión de Bash")
	return ""
}
