package controller

import (
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
	id, err := model.IDFromHeader(r.Header)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session := model.NewSession(id)
	model.AddSession(session)

	w.WriteHeader(http.StatusCreated)

	sessionJSON := model.MakeSessionJSON(session)
	sessionJSON.EncodeTo(w)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := model.SessionFromHeader(r.Header)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	model.RemoveSession(session)
	w.WriteHeader(http.StatusOK)
}
