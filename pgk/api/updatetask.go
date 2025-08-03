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
		writeJson(w, map[string]string{"error": "Ошибка десериализации JSON"})
		return
	}

	// Проверяем обязательное поле title
	if task.Title == "" {
		writeJson(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверяем обязательное поле id
	if task.ID == "" {
		writeJson(w, map[string]string{"error": "Не указан идентификатор задачи"})
		return
	}

	// Проверяем и обрабатываем дату
	if err := checkDate(&task); err != nil {
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	// Обновляем задачу в базе данных
	err := db.UpdateTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	// Возвращаем пустой JSON объект
	writeJson(w, map[string]interface{}{})
}
