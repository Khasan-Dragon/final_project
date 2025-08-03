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
		writeJson(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Получаем задачу из базы данных
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	// Проверяем, есть ли правило повторения
	if task.Repeat == "" {
		// Если правила повторения нет, удаляем задачу
		err = db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error": "Ошибка удаления задачи"})
			return
		}
	} else {
		// Если есть правило повторения, вычисляем следующую дату
		now := time.Now()
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJson(w, map[string]string{"error": "Ошибка вычисления следующей даты"})
			return
		}

		// Обновляем дату задачи
		err = db.UpdateDate(nextDate, id)
		if err != nil {
			writeJson(w, map[string]string{"error": "Ошибка обновления даты задачи"})
			return
		}
	}

	// Возвращаем пустой JSON объект
	writeJson(w, map[string]interface{}{})
}
