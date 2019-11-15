package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dgravesa/SessionServer/model"
)

type sessionValidResponse struct {
	Valid bool `json:"isValid"`
}

func makeSessionValidResponse(session model.Session) sessionValidResponse {
	return sessionValidResponse{
		Valid: model.IsValid(session),
	}
}

func validFunc(w http.ResponseWriter, r *http.Request) {
	const maxBodySize = 2048
	limitBody := io.LimitReader(r.Body, maxBodySize)
	session, err := model.ParseSession(limitBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := makeSessionValidResponse(session)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
