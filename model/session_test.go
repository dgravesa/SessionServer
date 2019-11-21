package model

import (
	"fmt"
	"net/http"
	"testing"
)

func validateSession(s Session, expectedUserID uint64) error {
	if s.UserID != expectedUserID {
		return fmt.Errorf("expected UserID = %d, actual UserID = %d", expectedUserID, s.UserID)
	} else if err := keycheck(s.Key); err != nil {
		return fmt.Errorf("%s", err)
	} else if err := keycheck(s.KeyHash); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func Test_NewSession_ReturnsValidSession(t *testing.T) {
	// Arrange
	var inputUserID uint64 = 12345

	// Act
	session := NewSession(inputUserID)

	// Assert
	if err := validateSession(session, inputUserID); err != nil {
		t.Error(err)
	}
}

func Test_SessionFromHeader_ReturnsError_OnMissingId(t *testing.T) {
	// Arrange
	h := http.Header{}
	h.Set("key", keygen())

	// Act
	_, err := SessionFromHeader(h)

	// Assert
	if err == nil {
		t.Errorf("missing ID in header did not produce an error")
	}
}

func Test_SessionFromHeader_ReturnsError_OnMissingKey(t *testing.T) {
	// Arrange
	h := http.Header{}
	h.Set("id", "123")

	// Act
	_, err := SessionFromHeader(h)

	// Assert
	if err == nil {
		t.Errorf("missing key in header did not produce an error")
	}
}

func Test_SessionFromHeader_ReturnsNoError_OnValidInput(t *testing.T) {
	// Arrange
	h := http.Header{}
	h.Set("id", "123")
	h.Set("key", keygen())

	// Act
	_, err := SessionFromHeader(h)

	if err != nil {
		t.Errorf("expected valid session, returned error: %s", err)
	}
}

func Test_SessionFromHeader_ReturnsValidSession_OnValidInput(t *testing.T) {
	// Arrange
	h := http.Header{}
	h.Set("id", "123")
	h.Set("key", keygen())

	// Act
	s, _ := SessionFromHeader(h)

	if err := validateSession(s, 123); err != nil {
		t.Errorf("expected valid session, returned error: %s", err)
	}
}
