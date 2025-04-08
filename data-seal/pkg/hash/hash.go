package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

func CalculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CalculateDirectoryHash(dirPath string) (string, error) {
	directoryHash := sha256.New()
	err := filepath.Walk(dirPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		fileHash, err := CalculateFileHash(filePath)
		if err != nil {
			return err
		}

		_, err = directoryHash.Write([]byte(fileHash))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(directoryHash.Sum(nil)), nil
}

func CalculateFinalHashFromIntegrityDir(hashesDir string) (string, error) {
	finalHash := sha256.New()

	err := filepath.Walk(hashesDir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		_, err = finalHash.Write(fileContent)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(finalHash.Sum(nil)), nil
}
