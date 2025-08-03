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
		writeJson(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Удаляем задачу из базы данных
	err := db.DeleteTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	// Возвращаем пустой JSON объект
	writeJson(w, map[string]interface{}{})
}
