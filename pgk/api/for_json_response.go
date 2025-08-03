package api

import (
	"encoding/json"
	"net/http"
)

// writeJson сериализует data в JSON и отправляет ответ
func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// writeErrorJson отправляет JSON ошибку с указанным HTTP кодом
func writeErrorJson(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
