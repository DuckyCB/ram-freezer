package constants

import (
	"crypto/md5"
	"crypto/sha256"
	"hash"
)

var Hashes = map[string]func() hash.Hash{
	"md5":    md5.New,
	"sha256": sha256.New,
}

type SystemFile int

const (
	DataSeal SystemFile = iota
	GhostKeyboard
	ProjectManager
	RamScraper
)

var SystemFiles = map[SystemFile]string{
	DataSeal:       "/opt/ram-freezer/bin/data-seal",
	GhostKeyboard:  "/opt/ram-freezer/bin/ghost-keyboard",
	ProjectManager: "/opt/ram-freezer/bin/project-manager",
	RamScraper:     "/opt/ram-freezer/bin/ram-scraper",
}
