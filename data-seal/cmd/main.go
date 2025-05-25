package main

import (
	"data-seal/internal/logs"
	"data-seal/pkg/hash"
	"data-seal/utils/constants"
	"flag"
	"fmt"
	"path/filepath"
	"ram-freezer/utils/rfutils/pkg/rfutils"
	"sync"
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
		hash.Checksum(outPath)
	}

	if *systemPtr {
		logs.Log.Info("Creating system hash")
		var wg sync.WaitGroup

		for _, filePath := range constants.SystemFiles {
			wg.Add(1)
			go func(filePath string) {
				defer wg.Done()
				hash.File(filePath)
			}(filePath)
		}

		// Installation
		version := rfutils.GetVersion()
		versionPath := filepath.Join("/opt/ram-freezer/bin/install", version)
		hash.Dir(versionPath)
		// Scripts
		hash.Dir("/opt/ram-freezer/bin/scripts")

	}
}
