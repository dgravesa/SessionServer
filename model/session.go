package model

import (
	"fmt"
	"net/http"
	"strconv"
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

// IDFromHeader gets the "id" from an HTTP header's query parameters.
func IDFromHeader(h http.Header) (uint64, error) {
	idStr := h.Get("id")
	return strconv.ParseUint(idStr, 10, 64)
}

// SessionKeyFromHeader gets the session "key" from an HTTP header's query parameters.
func SessionKeyFromHeader(h http.Header) (string, error) {
	var key string

	if key = h.Get("key"); keycheck(key) != nil {
		return "", fmt.Errorf("session key: invalid format on query parameter")
	}

	return key, nil
}

// SessionFromHeader constructs a session from an HTTP header's query parameters.
func SessionFromHeader(h http.Header) (Session, error) {
	var s Session
	var err error

	if s.UserID, err = IDFromHeader(h); err != nil {
		return s, err
	} else if s.Key, err = SessionKeyFromHeader(h); err != nil {
		return s, err
	}

	s.KeyHash = hashkey(s.Key)
	return s, nil
}
