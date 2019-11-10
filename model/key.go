package model

import (
	"crypto/rand"
	"encoding/hex"
)

const keyLength = 32

func keygen() string {
	keyBytes := make([]byte, keyLength)
	rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}

func keycheck(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
