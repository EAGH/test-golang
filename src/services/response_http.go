package services

import (
	"encoding/json"
	"gotest/src/models"
	"net/http"
)

// SendJSONResponse return a path of binary
func SendJSONResponse(msg string, status int, cont interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	response := models.MessageResponseHTTP{
		Status:  status,
		Message: msg,
		Content: cont,
	}
	json.NewEncoder(w).Encode(response)
}
