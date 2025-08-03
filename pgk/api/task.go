package api

import (
	"net/http"
)

// taskHandler обрабатывает запросы к /api/task в зависимости от HTTP метода
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		// Для других методов пока возвращаем ошибку
		writeErrorJson(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
