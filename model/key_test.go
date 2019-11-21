package model

import "testing"

func Test_keycheck_ReturnsNil_OnValidKey(t *testing.T) {
	// Arrange
	validKey := "ABCD1234567890FEDCBAABCD1234567890FEDCBAABCD1234567890FEDCBAAAAA"

	// Act
	err := keycheck(validKey)

	// Assert
	if err != nil {
		t.Errorf("expected key = \"%s\" to pass; returned error = \"%s\"", validKey, err)
	}
}

func Test_keycheck_ReturnsError_OnInvalidHex(t *testing.T) {
	// Arrange
	invalidKey := "ThisIsNotValidHex"

	// Act
	err := keycheck(invalidKey)

	// Assert
	if err == nil {
		t.Errorf("expected key = \"%s\" to fail; returned error = nil", invalidKey)
	}
}

func Test_keycheck_ReturnsError_OnInvalidLength(t *testing.T) {
	// Arrange
	invalidKey := "ABCD1234567890FEDCBAABCD12345678"

	// Act
	err := keycheck(invalidKey)

	// Assert
	if err == nil {
		t.Errorf("expected key = \"%s\" to fail; returned error = nil", invalidKey)
	}
}

func Test_keygen_ReturnsValidKey(t *testing.T) {
	// Arrange
	key := keygen()

	// Act
	err := keycheck(key)

	// Assert
	if err != nil {
		t.Errorf("expected generated key to pass check; keygen() = \"%s\", keycheck() = \"%s\"", key, err)
	}
}
