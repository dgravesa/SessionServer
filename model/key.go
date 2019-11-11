package model

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const keyLength = 32

func hashkey(key string) string {
	keyBytes, _ := hex.DecodeString(key)
	hashBytes := sha256.Sum256(keyBytes)
	return hex.EncodeToString(hashBytes[:])
}

func keygen() string {
	keyBytes := make([]byte, keyLength)
	rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}

func keycheck(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
