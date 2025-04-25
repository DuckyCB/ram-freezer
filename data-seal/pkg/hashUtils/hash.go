package hashUtils

import (
	"crypto/sha256"
	"data-seal/internal/logs"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"hash"
	"strings"
	"data-seal/utils/constants"
)

func CalculateFileHash(filePath string, hash hash.Hash) (string, error) { // hash es un puntero a sha256.New() o sha256.New() hash := sha256.New()
	file, err := os.Open(filePath)
	if err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(hash, file); err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}

	fileHash := hex.EncodeToString(hash.Sum(nil))

	return fileHash, nil
}

func CalculateDirectoryHash(dirPath string) (string, error) {
	var files []string

	// Recorre el directorio y guarda todos los paths de archivos
	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			logs.Log.Error(err.Error())
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}

	// Ordena los paths para asegurar un hash determinista
	sort.Strings(files)

	hash := sha256.New()

	for _, filePath := range files {
		fileHash, err := CalculateFileHash(filePath, hash)
		if err != nil {
			logs.Log.Error(fmt.Sprintf("Error al calcular el hash de %s: %v", filePath, err))
			return "", err
		}

		// Escribe tanto el nombre del archivo como su hash
		hash.Write([]byte(filePath))
		hash.Write([]byte(fileHash))
	}

	dirHash := hex.EncodeToString(hash.Sum(nil))
	logs.Log.Info(fmt.Sprintf("El hash del directorio %s es %s", dirPath, dirHash))

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


func IsHashFile(path string) bool {
	for key := range constants.Hashes {
		if strings.HasSuffix(path, key) {
			// Si el archivo ya tiene un hash, no lo procesamos
			return true
		}
	}
	return false
}