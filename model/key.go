package model

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func keygen() string {
	keyBytes := make([]byte, sha256.Size)
	rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}

func keycheck(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
