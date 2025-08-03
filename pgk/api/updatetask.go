package api

import (
	"encoding/json"
	"final_project/pgk/db"
	"net/http"
)

// updateTaskHandler обрабатывает PUT запросы для обновления задач
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Десериализуем JSON в структуру Task
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeErrorJson(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	// Проверяем обязательное поле title
	if task.Title == "" {
		writeErrorJson(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	// Проверяем обязательное поле id
	if task.ID == "" {
		writeErrorJson(w, "Не указан идентификатор задачи", http.StatusBadRequest)
		return
	}

	// Проверяем и обрабатываем дату
	if err := checkDate(&task); err != nil {
		writeErrorJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем задачу в базе данных
	err := db.UpdateTask(&task)
	if err != nil {
		writeErrorJson(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	// Возвращаем пустой JSON объект
	writeJson(w, map[string]interface{}{})
}
