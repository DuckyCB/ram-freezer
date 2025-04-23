package hash

import (
	"crypto/sha256"
	"data-seal/internal/logs"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CalculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}

	fileHash := hex.EncodeToString(hash.Sum(nil))
	logs.Log.Info(fmt.Sprintf("el hash de %s es %s", filePath, fileHash))

	return fileHash, nil
}

func CalculateDirectoryHash(dirPath string) (string, error) {
	hash := sha256.New()

	err := filepath.Walk(dirPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		fileHash, err := CalculateFileHash(filePath)
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}

		_, err = hash.Write([]byte(fileHash))
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}

	dirHash := hex.EncodeToString(hash.Sum(nil))
	logs.Log.Info(fmt.Sprintf("el hash de %s es %s", dirPath, dirHash))

	return dirHash, nil
}

func CalculateFinalHashFromIntegrityDir(hashesDir string) (string, error) {
	hash := sha256.New()

	err := filepath.Walk(hashesDir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}

		_, err = hash.Write(fileContent)
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}

	calculatedHash := hex.EncodeToString(hash.Sum(nil))
	logs.Log.Info(fmt.Sprintf("el hash de %s es %s", hashesDir, calculatedHash))

	return calculatedHash, nil
}
