package api

import (
	"final_project/pgk/db"
	"net/http"
)

// getTaskHandler обрабатывает GET запросы для получения задачи по ID
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр id из запроса
	id := r.URL.Query().Get("id")
	if id == "" {
		writeErrorJson(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	// Получаем задачу из базы данных
	task, err := db.GetTask(id)
	if err != nil {
		writeErrorJson(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	// Возвращаем задачу в JSON формате
	writeJson(w, task)
}
