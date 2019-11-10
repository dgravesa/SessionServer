package model

// DataLayer is an interface to the Session Server data store.
type DataLayer interface {
	AddSession(s *Session)
	RemoveSession(s *Session)
	IsValid(s *Session) bool
}

var dataLayer DataLayer

// SetData injects layer as the data store.
func SetData(layer DataLayer) {
	dataLayer = layer
}

// AddSession adds a user session to the data store.
func AddSession(s *Session) {
	dataLayer.AddSession(s)
}

// RemoveSession removes a user session from the data store if it exists.
func RemoveSession(s *Session) {
	dataLayer.RemoveSession(s)
}

// IsValid returns true if the session is found in the data store; false otherwise.
func IsValid(s *Session) bool {
	return dataLayer.IsValid(s)
}
