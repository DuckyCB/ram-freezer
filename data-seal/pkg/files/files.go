package files

import (
	"data-seal/internal/logs"
	"fmt"
	"os"
	"ram-freezer/utils/rfutils/pkg/rfutils"
	"sync"
)

var sha256Mutex = sync.Mutex{}
var md5Mutex = sync.Mutex{}
var hashFilePath string

func WriteChecksum(data, hash string) error {
	checksumPath := fmt.Sprintf("%s/CHECKSUM.%s", rfutils.GetOutPath(), hash)

	err := os.WriteFile(checksumPath, []byte(data), 0644)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	logs.Log.Info(fmt.Sprintf("checksum (%s) %s guardado en %s", data, hash, checksumPath))
	return nil
}

func WriteToFile(name, data, hash string) error {
	if hashFilePath == "" {
		hashFilePath = fmt.Sprintf("%s/file-hashes", rfutils.GetOutPath())
		logs.Log.Info(fmt.Sprintf("%s configurado como base de hashes", hashFilePath))
	}

	if hash == "sha256" {
		err := writeToSha256File(hashFilePath, fmt.Sprintf("%s: %s\n", name, data))
		if err != nil {
			return err
		}
	} else {
		err := writeToMD5File(hashFilePath, fmt.Sprintf("%s: %s\n", name, data))
		if err != nil {
			return err
		}
	}

	return nil
}

func writeToSha256File(filePath, data string) error {
	sha256Mutex.Lock()
	defer sha256Mutex.Unlock()

	sha256Path := fmt.Sprintf("%s.sha256", filePath)

	err := os.WriteFile(sha256Path, []byte(data), 0644)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	logs.Log.Info(fmt.Sprintf("sha256 de %s guardado en %s", filePath, sha256Path))
	return nil
}

func writeToMD5File(filePath, data string) error {
	md5Mutex.Lock()
	defer md5Mutex.Unlock()

	md5Path := fmt.Sprintf("%s.md5", filePath)

	err := os.WriteFile(md5Path, []byte(data), 0644)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	logs.Log.Info(fmt.Sprintf("md5 de %s guardado en %s", filePath, md5Path))
	return nil
}
