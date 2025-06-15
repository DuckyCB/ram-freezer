package main

import (
	"data-seal/internal/logs"
	"data-seal/internal/utils/constants"
	"data-seal/pkg/hash"
	"flag"
	"fmt"
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
		logs.Log.Info(fmt.Sprintf("Hashing file: %s", *filePtr))
		hash.File(*filePtr)
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

		// System info
		hash.File(outPath + "/system.info")
		// Installation
		version := rfutils.GetVersion()
		installPath := fmt.Sprintf("/opt/ram-freezer/bin/install/%s.log", version)
		hash.File(installPath)
		// Scripts
		hash.Dir("/opt/ram-freezer/bin/scripts")
	}

	if *checksumPtr {
		logs.Log.Info("Creating checksum")
		hash.Checksum(outPath)
	}

	logs.Log.Info("Exiting data-seal")
}
