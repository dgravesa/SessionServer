package model

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	keyLength    = 32
	keyLengthHex = keyLength * 2
)

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

func keycheck(s string) error {
	if _, err := hex.DecodeString(s); err != nil {
		return fmt.Errorf("invalid key string format: %s", s)
	}

	if len(s) != keyLengthHex {
		return fmt.Errorf("invalid hash string length: expected = %d, actual = %d", keyLengthHex, len(s))
	}

	return nil
}
