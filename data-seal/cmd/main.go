package main

import (
	"data-seal/internal/logs"
	"data-seal/pkg/files"
	"data-seal/pkg/hash"
	"flag"
	"fmt"
	"path/filepath"
	"ram-freezer/utils/rfutils/pkg/rfutils"
)

func main() {
	logs.SetupLogger()

	dirPtr := flag.String("dir", "", "directory to hash")
	filePtr := flag.String("file", "", "file to hash")
	checksumPtr := flag.Bool("checksum", false, "create checksum file")
	systemPtr := flag.Bool("system", false, "create a hash of the system")
	flag.Parse()

	logs.Log.Info("Starting data-seal")

	outPath := rfutils.GetOutPath()

	if *dirPtr != "" {
		logs.Log.Info(fmt.Sprintf("Hashing directory: %s", *dirPtr))
		hash.Dir(*dirPtr)
	}

	if *filePtr != "" {
		hash.File(*filePtr)
	}

	if *checksumPtr {
		logs.Log.Info("Creating checksum")

		checksumHash, err := hash.Checksum(outPath)
		if err != nil {
			logs.Log.Error(err.Error())
			return
		}

		logs.Log.Info(fmt.Sprintf("El hash encadenado es %s", checksumHash))

		err = files.WriteToFile(filepath.Join(outPath, "CHECKSUM"), checksumHash)
	}

	if *systemPtr {
		logs.Log.Info("Creating system hash")
	}
}
