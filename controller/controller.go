package controller

import (
	"net/http"
)

// RegisterAll initializes the route handlers for the session server.
func RegisterAll() {
	http.HandleFunc("/session", sessionFunc)
	http.HandleFunc("/sessionValid", validFunc)
}
