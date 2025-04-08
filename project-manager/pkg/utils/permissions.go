package utils

import (
	"os"
	"os/exec"
	"strings"
)

func IsAdmin() bool {
	if os.Getuid() == 0 {
		return true
	}

	cmd := exec.Command("groups")
	outputBytes, err := cmd.Output()
	if err == nil {
		groups := strings.Fields(string(outputBytes))
		for _, group := range groups {
			if group == "sudo" || group == "wheel" {
				return true
			}
		}
	}

	return false
}
