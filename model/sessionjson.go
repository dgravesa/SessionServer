package model

import (
	"encoding/json"
	"fmt"
	"io"
)

// SessionJSON contains information from a Session to be passed around in JSON format.
type SessionJSON struct {
	UserID uint64 `json:"userId"`
	Key    string `json:"key"`
}

// MakeSessionJSON creates a new SessionJSON from a Session.
func MakeSessionJSON(s Session) SessionJSON {
	var sj SessionJSON
	sj.FromSession(s)
	return sj
}

// ToSession creates a Session corresponding to the SessionJSON.
func (sj SessionJSON) ToSession() Session {
	var session Session
	session.UserID = sj.UserID
	session.Key = sj.Key
	session.KeyHash = hashkey(sj.Key)
	return session
}

// FromSession pulls the SessionJSON from a Session.
func (sj *SessionJSON) FromSession(s Session) {
	sj.UserID = s.UserID
	sj.Key = s.Key
}

// EncodeTo writes the SessionJSON data in JSON format to a Writer.
func (sj SessionJSON) EncodeTo(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(sj)
}

// DecodeFrom gets the SessionJSON data from a Reader.
func (sj *SessionJSON) DecodeFrom(r io.Reader) error {
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

	sj.UserID = *nillable.UserID
	sj.Key = *nillable.Key

	return nil
}
