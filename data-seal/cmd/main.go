package main

import (
	"data-seal/pkg/files"
	"data-seal/pkg/hash"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

var basePath = "/opt/ram-freezer/bin/"
var hashesDir = "integrity"

func main() {
	dirPtr := flag.String("dir", "", "directory to hash")
	filePtr := flag.String("file", "", "file to hash")
	//allPtr := flag.Bool("all", false, "hash all files")
	//finalPtr := flag.Bool("final", false, "hash all files")
	chainPtr := flag.Bool("chain", false, "create a hash chain")
	flag.Parse()

	log.Println("starting data seal")
	log.Printf("dir: %s, file: %s, chain: %t", *dirPtr, *filePtr, *chainPtr)

	if *dirPtr != "" {
		log.Printf("hashing dir: %s", *dirPtr)
		dirPath := filepath.Join(basePath, *dirPtr)
		dirHash, err := hash.CalculateDirectoryHash(dirPath)
		if err != nil {
			fmt.Printf("Error calculating hash for %s: %v\n", *dirPtr, err)
			return
		}
		dirName := filepath.Base(dirPath)

		log.Printf("El hash de %s es %s", dirName, dirHash)

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, dirName), dirHash)
	}

	if *filePtr != "" {
		log.Printf("hashing file: %s", *filePtr)
		filePath := filepath.Join(basePath, *filePtr)
		fileHash, err := hash.CalculateFileHash(filePath)
		if err != nil {
			fmt.Printf("Error calculating hash for %s: %v\n", *filePtr, err)
			return
		}
		fileName := filepath.Base(filePath)

		log.Printf("El hash de %s es %s", fileName, fileHash)

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, fileName), fileHash)
	}

	if *chainPtr {
		log.Println("creating hash chain")
		chainHash, err := hash.CalculateFinalHashFromIntegrityDir(filepath.Join(basePath, hashesDir))
		if err != nil {
			fmt.Println("Error calculating chain hash:", err)
			return
		}

		log.Printf("El hash encadenado es %s", chainHash)

		err = files.WriteToFile(filepath.Join(basePath, hashesDir, "chain"), chainHash)
	}
}
