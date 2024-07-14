package utils

import (
	"encoding/json"
	"net/http"
)

// respondWithJSON serializa a resposta em JSON e define os headers apropriados.
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}