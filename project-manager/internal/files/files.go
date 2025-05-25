package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"project-manager/internal/logs"
	"ram-freezer/utils/rfutils/pkg/rfutils"
)

// Helper function to copy a file
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	logs.Log.Info(fmt.Sprintf("%s copiado a %s", in.Name(), out.Name()))
	return nil
}

// Helper function to copy a directory recursively
func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	if !srcInfo.IsDir() {
		logs.Log.Error(fmt.Sprintf("%s is not a directory", src))
		return fmt.Errorf("%s is not a directory", src)
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				logs.Log.Error(err.Error())
				continue
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				logs.Log.Warn(fmt.Sprintf("El archivo %s no fue copiado", srcPath))
				continue
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
	copyFile(fmt.Sprintf("%s/install/%s.log", sourceBinDir, version), fmt.Sprintf("%s/%s.log", destinationDirOnUSB, version))

	// Run
	copyDir(runPath, destinationDirOnUSB)

	// Scripts
	copyDir(filepath.Join(sourceBinDir, "scripts"), destinationDirOnUSB)

	// Binaries
	copyFile(filepath.Join(sourceBinDir, "data-seal"), filepath.Join(destinationDirOnUSB, "data-seal"))
	copyFile(filepath.Join(sourceBinDir, "ghost-keyboard"), filepath.Join(destinationDirOnUSB, "ghost-keyboard"))
	copyFile(filepath.Join(sourceBinDir, "project-manager"), filepath.Join(destinationDirOnUSB, "project-manager"))
	copyFile(filepath.Join(sourceBinDir, "ram-scraper"), filepath.Join(destinationDirOnUSB, "ram-scraper"))

	// TODO: falta mover archivos que ya estan en el USB a la nueva ubicacion
}
