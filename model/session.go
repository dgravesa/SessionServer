package model

import (
	"io"
)

// Session maintains a user login instance.
type Session struct {
	UserID  uint64
	Key     string
	KeyHash string
}

// NewSession creates a user session with a randomly generated key.
func NewSession(uid uint64) Session {
	var s Session
	s.UserID = uid
	s.Key = keygen()
	s.KeyHash = hashkey(s.Key)
	return s
}

// ParseSession constructs a session from an HTTP request.
func ParseSession(r io.Reader) (Session, error) {
	var s Session

	var sj SessionJSON
	if err := sj.DecodeFrom(r); err != nil {
		return s, err
	}

	s = sj.ToSession()
	return s, nil
}
