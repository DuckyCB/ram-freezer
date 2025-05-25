package hash

import (
	"data-seal/internal/logs"
	"data-seal/pkg/files"
	"data-seal/utils/constants"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func hashFile(filePath string, hash hash.Hash) (string, error) {
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

func File(filePath string) error {
	if filePath == "" {
		logs.Log.Error("File path is empty")
		return fmt.Errorf("file is empty")
	}

	logs.Log.Info(fmt.Sprintf("Hashing file: %s", filePath))

	for hashName, hashObj := range constants.Hashes {
		logs.Log.Info(fmt.Sprintf("Calculando %s para %s", hashName, filePath))

		hashValue, err := hashFile(filePath, hashObj())

		if err == nil {
			logs.Log.Info(fmt.Sprintf("Hash %s para %s: %s", hashName, filePath, hashValue))
		} else {
			logs.Log.Error(fmt.Sprintf("Error calculando %s para %s: %v", hashName, filePath, err))
			continue
		}

		err = files.WriteToFile(filepath.Base(filePath), hashValue, hashName)
		if err != nil {
			return err
		}
	}
	return nil
}

func Dir(dirPath string) {
	var wg sync.WaitGroup

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			logs.Log.Error(err.Error())
			return nil
		}

		if path == dirPath {
			return nil
		}

		if IsHashFile(path) {
			logs.Log.Info(fmt.Sprintf("Borrando archivo hash: %s", path))
			err := os.Remove(path)
			if err != nil {
				logs.Log.Error(fmt.Sprintf("Error borrando archivo hash: %s", err.Error()))
				return nil
			}
			return nil
		}

		if d.IsDir() {
			Dir(path)
		} else if !d.IsDir() {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				File(p)
			}(path)
		}
		return nil
	})
	if err != nil {
		logs.Log.Error(err.Error())
		return
	}

	wg.Wait()
	return
}

func Checksum(outPath string) {
	for hashName, hashFunc := range constants.Hashes {
		hashValue, err := hashFile(fmt.Sprintf("%s/file-hashes.%s", outPath, hashName), hashFunc())

		if err == nil {
			logs.Log.Info(fmt.Sprintf("CHECKSUM %s: %s", hashName, hashValue))
		} else {
			logs.Log.Error(fmt.Sprintf("Error calculando CHECKSUM %s", hashName))
			continue
		}

		files.WriteChecksum(hashValue, hashName)
	}
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
