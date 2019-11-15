package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dgravesa/SessionServer/model"
)

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

	sessionJSON := model.MakeSessionJSON(session)
	sessionJSON.EncodeTo(w)
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
