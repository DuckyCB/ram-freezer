package main

import (
	"data-seal/internal/logs"
	"data-seal/pkg/files"
	"data-seal/pkg/hash"
	"flag"
	"fmt"
	"path/filepath"
)

var basePath = "/opt/ram-freezer/bin/"
var hashesDir = "integrity"

func main() {
	logs.SetupLogger()

	dirPtr := flag.String("dir", "", "directory to hash")
	filePtr := flag.String("file", "", "file to hash")
	//allPtr := flag.Bool("all", false, "hash all files")
	//finalPtr := flag.Bool("final", false, "hash all files")
	chainPtr := flag.Bool("chain", false, "create a hash chain")
	flag.Parse()

	if *dirPtr != "" {
		logs.Log.Info(fmt.Sprintf("hashing dir: %s", *dirPtr))

		dirPath := filepath.Join(basePath, *dirPtr)
		dirHash, err := hash.CalculateDirectoryHash(dirPath)
		if err != nil {
			logs.Log.Error(fmt.Sprintf("Error calculating hash for %s: %v\n", *dirPtr, err))
			return
		}
		dirName := filepath.Base(dirPath)

		logs.Log.Info(fmt.Sprintf("El hash de %s es %s", dirName, dirHash))

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, dirName), dirHash)
	}

	if *filePtr != "" {
		logs.Log.Info(fmt.Sprintf("hashing file: %s", *filePtr))

		filePath := filepath.Join(basePath, *filePtr)
		fileHash, err := hash.CalculateFileHash(filePath)
		if err != nil {
			logs.Log.Error(fmt.Sprintf("Error calculating hash for %s: %v\n", *filePtr, err))
			return
		}
		fileName := filepath.Base(filePath)

		logs.Log.Info(fmt.Sprintf("El hash de %s es %s", fileName, fileHash))

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, fileName), fileHash)
	}

	if *chainPtr {
		logs.Log.Info("Creating hash chain")

		chainHash, err := hash.CalculateFinalHashFromIntegrityDir(filepath.Join(basePath, hashesDir))
		if err != nil {
			logs.Log.Error(err.Error())
			return
		}

		logs.Log.Info(fmt.Sprintf("El hash encadenado es %s", chainHash))

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, "chain"), chainHash)
	}
}
