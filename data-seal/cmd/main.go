package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var basePath = "/opt/ram-freezer/bin/"
var hashesDir = "integrity"

func calculateFileHash(filePath string) (string, error) {
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

func calculateDirectoryHash(dirPath string) (string, error) {
	directoryHash := sha256.New()
	err := filepath.Walk(dirPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		fileHash, err := calculateFileHash(filePath)
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

func calculateFinalHashFromIntegrityDir() (string, error) {
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

func writeToFile(filePath, data string) error {
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return err
	}
	return nil
}

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
		dirHash, err := calculateDirectoryHash(dirPath)
		if err != nil {
			fmt.Printf("Error calculating hash for %s: %v\n", *dirPtr, err)
			return
		}
		dirName := filepath.Base(dirPath)

		log.Printf("El hash de %s es %s", dirName, dirHash)

		err = writeToFile(filepath.Join(basePath, hashesDir, dirName), dirHash)
	}

	if *filePtr != "" {
		log.Printf("hashing file: %s", *filePtr)
		filePath := filepath.Join(basePath, *filePtr)
		fileHash, err := calculateFileHash(filePath)
		if err != nil {
			fmt.Printf("Error calculating hash for %s: %v\n", *filePtr, err)
			return
		}
		fileName := filepath.Base(filePath)

		log.Printf("El hash de %s es %s", fileName, fileHash)

		err = writeToFile(filepath.Join(basePath, hashesDir, fileName), fileHash)
	}

	if *chainPtr {
		log.Println("creating hash chain")
		chainHash, err := calculateFinalHashFromIntegrityDir()
		if err != nil {
			fmt.Println("Error calculating chain hash:", err)
			return
		}

		log.Printf("El hash encadenado es %s", chainHash)

		err = writeToFile(filepath.Join(basePath, hashesDir, "chain"), chainHash)
	}
}
