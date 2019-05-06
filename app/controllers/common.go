package controllers

import (
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		throwInternalServerError(w, err)
		return
	}

	body := make(map[string]interface{})

	body["code"] = status
	body["message"] = response
	bodyData, err := json.Marshal(body)

	if err != nil {
		throwInternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, _ = w.Write([]byte(bodyData))
}

func throwInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
