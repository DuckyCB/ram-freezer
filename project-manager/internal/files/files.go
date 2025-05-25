package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"project-manager/internal/logs"
	"ram-freezer/utils/rfutils/pkg/rfutils"
)

// Helper function to copy a file
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("failed to copy file from %s to %s: %w", src, dst, err)
	}
	return nil
}

// Helper function to copy a directory recursively
func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to get source directory info %s: %w", src, err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("source %s is not a directory", src)
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to create destination directory %s: %w", dst, err)
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read source directory %s: %w", src, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyToUSB() {
	sourceBinDir := "/opt/ram-freezer/bin"
	usbMountPoint := "/mnt/usb"

	runPath := rfutils.GetOutPath()
	version := rfutils.GetVersion()
	runName := filepath.Base(runPath)
	destinationDirOnUSB := filepath.Join(usbMountPoint, runName)

	logs.Log.Info("Creating destination directory: " + destinationDirOnUSB)
	if err := os.MkdirAll(destinationDirOnUSB, 0755); err != nil {
		logs.Log.Error(err.Error())
	}
	logs.Log.Info("Destination directory created: " + destinationDirOnUSB)

	// Installation
	installationPath := fmt.Sprintf("%s/install/%s.log", sourceBinDir, version)
	installationNewPath := fmt.Sprintf("%s/%s.log", destinationDirOnUSB, version)
	logs.Log.Info(fmt.Sprintf("Attempting to copy %s to %s", installationPath, installationPath))
	if err := copyFile(installationPath, installationNewPath); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copied %s to %s", installationPath, installationNewPath))
	}

	// Run
	if err := copyDir(runPath, destinationDirOnUSB); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copies run files to %s", destinationDirOnUSB))
	}

	// Scripts
	if err := copyDir(filepath.Join(sourceBinDir, "scripts"), destinationDirOnUSB); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copied scripts to %s", destinationDirOnUSB))
	}

	// Binaries
	if err := copyFile(filepath.Join(sourceBinDir, "data-seal"), filepath.Join(destinationDirOnUSB, "data-seal")); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copied data-seal to %s", destinationDirOnUSB))
	}

	if err := copyFile(filepath.Join(sourceBinDir, "ghost-keyboard"), filepath.Join(destinationDirOnUSB, "ghost-keyboard")); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copied ghost-keyboard to %s", destinationDirOnUSB))
	}

	if err := copyFile(filepath.Join(sourceBinDir, "project-manager"), filepath.Join(destinationDirOnUSB, "project-manager")); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copied project-manager to %s", destinationDirOnUSB))
	}

	if err := copyFile(filepath.Join(sourceBinDir, "ram-scraper"), filepath.Join(destinationDirOnUSB, "ram-scraper")); err != nil {
		logs.Log.Error(err.Error())
	} else {
		logs.Log.Info(fmt.Sprintf("Copied ram-scraper to %s", destinationDirOnUSB))
	}

	// TODO: falta mover archivos que ya estan en el USB a la nueva ubicacion
}
