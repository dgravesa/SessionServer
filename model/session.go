package model

import (
	"encoding/json"
	"fmt"
	"io"
)

// Session maintains a user login instance.
type Session struct {
	UserID uint64 `json:"userId"`
	Key    string `json:"key"`
}

// NewSession creates a user session with a randomly generated key.
func NewSession(uid uint64) *Session {
	s := new(Session)
	s.UserID = uid
	s.Key = keygen()
	return s
}

// ParseSession constructs a session from an HTTP request.
func ParseSession(r io.Reader) (*Session, error) {
	session := new(Session)
	err := decodeSessionJSON(r, session)
	return session, err
}

func decodeSessionJSON(r io.Reader, s *Session) error {
	var nillable struct {
		UserID *uint64 `json:"userId"`
		Key    *string `json:"key"`
	}

	d := json.NewDecoder(r)
	if err := d.Decode(&nillable); err != nil {
		return err
	} else if nillable.UserID == nil {
		return fmt.Errorf("Session: missing \"userId\" in JSON")
	} else if nillable.Key == nil {
		return fmt.Errorf("Session: missing \"key\" in JSON")
	} else if !keycheck(*nillable.Key) {
		return fmt.Errorf("Session: invalid key format")
	}

	s.UserID = *nillable.UserID
	s.Key = *nillable.Key

	return nil
}
