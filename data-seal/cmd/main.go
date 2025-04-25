package main

import (
	"data-seal/internal/logs"
	"data-seal/pkg/files"
	"data-seal/pkg/hashUtils"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"data-seal/utils/constants"
)

var basePath = "/opt/ram-freezer/bin/"
var hashesDir = "integrity"

// Mutex global para logs
var mu sync.Mutex


func main() {
	logs.SetupLogger()

	dirPtr := flag.String("dir", "", "directory to hash")
	filePtr := flag.String("file", "", "file to hash")
	//allPtr := flag.Bool("all", false, "hash all files")
	//finalPtr := flag.Bool("final", false, "hash all files")
	chainPtr := flag.Bool("chain", false, "create a hash chain")
	flag.Parse()

	logs.Log.Info("Starting data-seal")

	if *dirPtr != "" {
		logs.Log.Info(fmt.Sprintf("Hashing directory: %s", *dirPtr))
		hashDir(*dirPtr)
	}

	if *filePtr != "" {
		hashFile(*filePtr)
	}

	if *chainPtr {
		logs.Log.Info("Creating hash chain")

		chainHash, err := hashUtils.CalculateFinalHashFromIntegrityDir(filepath.Join(basePath, hashesDir))
		if err != nil {
			logs.Log.Error(err.Error())
			return
		}

		logs.Log.Info(fmt.Sprintf("El hash encadenado es %s", chainHash))

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, "chain"), chainHash)
	}
}

func hashDir(dirPath string) {
	var wg sync.WaitGroup

	// recorrer el directorio y calcular el hash de cada archivo
	_ = filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			mu.Lock()
			logs.Log.Error(err.Error())
			mu.Unlock()
			return err
		}
		
		if path == dirPath {
			return nil
		}

		if hashUtils.IsHashFile(path) {
			// Si el archivo es un archivo hash, lo borramos
			mu.Lock()
			logs.Log.Info(fmt.Sprintf("Borrando archivo hash: %s", path))
			mu.Unlock()
			err := os.Remove(path)
			if err != nil {
				mu.Lock()
				logs.Log.Error(fmt.Sprintf("Error borrando archivo hash: %s", err.Error()))
				mu.Unlock()
				return nil
			}
			return nil
		}

		if d.IsDir() {
			hashDir(path)
		} else if !d.IsDir() {
			// procesamiento paralelo de archivos
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				hashFile(p)
			}(path)
		}
		return nil
	})
	wg.Wait() // Espera a que terminen todas las tareas en paralelo
}

// Procesa archivos: calcula los hashes y los guarda en archivos
func hashFile(filePath string) {
	if filePath != "" {
		mu.Lock() // Bloqueamos acceso a logs
		logs.Log.Info(fmt.Sprintf("Hashing file: %s", filePath))
		mu.Unlock()

		for hashName, hashObj := range constants.Hashes {
			mu.Lock() // Bloqueamos acceso a logs
			logs.Log.Info(fmt.Sprintf("Calculando %s para %s", hashName, filePath))
			mu.Unlock()
			
			hashValue, err := hashUtils.CalculateFileHash(filePath, hashObj())

			mu.Lock()
			logs.Log.Info(fmt.Sprintf("Hash %s para %s: %s", hashName, filePath, hashValue))
			mu.Unlock()

			if err != nil {
				mu.Lock()
				logs.Log.Error(fmt.Sprintf("Error calculando %s para %s: %v", hashName, filePath, err))
				mu.Unlock()
				continue
			}

			hashFilePath := fmt.Sprintf("%s.%s", filePath, hashName)

			mu.Lock()
			logs.Log.Info(fmt.Sprintf("Escribiendo hash %s en %s", hashName, hashFilePath))
			mu.Unlock()

			err = files.WriteToFile(hashFilePath, hashValue)
			if err != nil {
				mu.Lock()
				logs.Log.Error(fmt.Sprintf("Error escribiendo archivo de hash %s: %v", hashName, err))
				mu.Unlock()
			}
		}
	}
}
