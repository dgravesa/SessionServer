package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"bitbucket.org/dangravesteam/WaterLogger/SessionServer/model"
)

// RegisterAll initializes the route handlers for the session server.
func RegisterAll() {
	http.HandleFunc("/session", sessionFunc)
	http.HandleFunc("/sessionValid", validFunc)
}

func sessionFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postSession(w, r)
	case http.MethodDelete:
		deleteSession(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postSession(w http.ResponseWriter, r *http.Request) {
	const maxBodySize = 512
	limitBody := io.LimitReader(r.Body, maxBodySize)
	decoder := json.NewDecoder(limitBody)

	var requestBody struct {
		UserID *uint64 `json:"userId"`
	}

	if err := decoder.Decode(&requestBody); err != nil || requestBody.UserID == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session := model.NewSession(*requestBody.UserID)
	model.AddSession(session)

	w.WriteHeader(http.StatusCreated)

	sessionJSON := model.MakeSessionJSON(*session)
	sessionJSON.Encode(w)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	const maxBodySize = 2048
	limitBody := io.LimitReader(r.Body, maxBodySize)
	session, err := model.ParseSession(limitBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	model.RemoveSession(session)
	w.WriteHeader(http.StatusOK)
}

type sessionValidResponse struct {
	Valid bool `json:"isValid"`
}

func makeSessionValidResponse(session *model.Session) *sessionValidResponse {
	return &sessionValidResponse{
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
