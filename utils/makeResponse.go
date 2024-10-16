package utils

import (
	"encoding/json"
	"net/http"
)

func MakeJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func MakeErrorResponse(w http.ResponseWriter, status int, errMessage string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": errMessage})
}
