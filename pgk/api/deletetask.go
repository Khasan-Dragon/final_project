package api

import (
	"final_project/pgk/db"
	"net/http"
)

// deleteTaskHandler обрабатывает DELETE запросы для удаления задач
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр id из запроса
	id := r.URL.Query().Get("id")
	if id == "" {
		writeErrorJson(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	// Удаляем задачу из базы данных
	err := db.DeleteTask(id)
	if err != nil {
		writeErrorJson(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	// Возвращаем пустой JSON объект
	writeJson(w, map[string]interface{}{})
}
