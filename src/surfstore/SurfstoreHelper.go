package surfstore

import (
	"crypto/sha256"
	"encoding/hex"
)

/* Hash Related helper function. You do not need to use or modify any of these functions */
func GetBlockHashBytes(blockData []byte) []byte {
	h := sha256.New()
	h.Write(blockData)
	return h.Sum(nil)
}

func GetBlockHashString(blockData []byte) string {
	blockHash := GetBlockHashBytes(blockData)
	return hex.EncodeToString(blockHash)
}
