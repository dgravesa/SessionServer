package model

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"text/template"
)

var jsonSessionTmpl *template.Template

type jsonSessionData struct {
	UserID uint64
	Key    string
}

func TestMain(m *testing.M) {
	jsonTmplStr := `{
		"userId": {{.UserID}},
		"key": "{{.Key}}"
	}`

	jsonSessionTmpl = template.Must(template.New("").Parse(jsonTmplStr))

	result := m.Run()

	os.Exit(result)
}

func validateHash(h string) error {
	if _, err := hex.DecodeString(h); err != nil {
		return fmt.Errorf("invalid hash string format: %s", h)
	}

	expectedLength := keyLength * 2 // 2 hex characters per byte

	if len(h) != expectedLength {
		return fmt.Errorf("invalid hash string length: expected = %d, actual = %d", expectedLength, len(h))
	}

	return nil
}

func validateSession(s Session, expectedUserID uint64) error {
	if s.UserID != expectedUserID {
		return fmt.Errorf("expected UserID = %d, actual UserID = %d", expectedUserID, s.UserID)
	} else if err := validateHash(s.Key); err != nil {
		return fmt.Errorf("%s", err)
	} else if err := validateHash(s.KeyHash); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func Test_NewSessionReturnsValidSession(t *testing.T) {
	// Arrange
	var inputUserID uint64 = 12345

	// Act
	session := NewSession(inputUserID)

	// Assert
	if err := validateSession(session, inputUserID); err != nil {
		t.Error(err)
	}
}

func createJSONSessionInputReader(userID uint64, key string) io.Reader {
	var jsonBuilder strings.Builder
	jsonSessionTmpl.Execute(&jsonBuilder, jsonSessionData{userID, key})

	return strings.NewReader(jsonBuilder.String())
}

func Test_ParseSessionReturnsValidSessionOnValidInput(t *testing.T) {
	// Arrange
	var inputUserID uint64 = 5678
	inputKey := "ABCD1234567890FEDCBAABCD1234567890FEDCBAABCD1234567890FEDCBAAAAA"
	r := createJSONSessionInputReader(inputUserID, inputKey)

	// Act
	session, err := ParseSession(r)

	// Assert
	if err != nil {
		t.Error(err)
	} else if err = validateSession(session, inputUserID); err != nil {
		t.Error(err)
	} else if session.Key != inputKey {
		t.Errorf("expected session key = [%s...%s], received session key = [%s...%s]",
			inputKey[:4], inputKey[len(inputKey)-4:], session.Key[:4], session.Key[len(session.Key)-4:])
	}
}
