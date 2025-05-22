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
