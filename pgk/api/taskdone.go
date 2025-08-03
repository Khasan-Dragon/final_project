package api

import (
	"final_project/pgk/db"
	"net/http"
	"time"
)

// taskDoneHandler обрабатывает POST запросы для выполнения задач
func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
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

	// Проверяем, есть ли правило повторения
	if task.Repeat == "" {
		// Если правила повторения нет, удаляем задачу
		err = db.DeleteTask(id)
		if err != nil {
			writeErrorJson(w, "Ошибка удаления задачи", http.StatusInternalServerError)
			return
		}
	} else {
		// Если есть правило повторения, вычисляем следующую дату
		now := time.Now()
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeErrorJson(w, "Ошибка вычисления следующей даты", http.StatusBadRequest)
			return
		}

		// Обновляем дату задачи
		err = db.UpdateDate(nextDate, id)
		if err != nil {
			writeErrorJson(w, "Ошибка обновления даты задачи", http.StatusInternalServerError)
			return
		}
	}

	// Возвращаем пустой JSON объект
	writeJson(w, map[string]interface{}{})
}
